package di

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mService struct{}

func (m *mService) HelloWord() (string, error) { return "hello", nil }

type (
	mockService1 struct{ m *mService }
	mockService2 struct{}
	mockService3 struct{}
	errService   struct{}
)

var (
	mockS1 = mockService1{m: &mService{}}
	mockS2 = mockService2{}
	mockS3 = mockService3{}
)

var ErrTestDestroy = fmt.Errorf("Hi")

//** Mock
func (ms1 *mockService1) Regist() error { return nil }
func (ms1 *mockService1) Destroy() error {
	ms1.m = nil
	return nil
}
func (ms1 *mockService1) Get() interface{} { return ms1.m }

func (ms *mockService2) Regist() error    { return nil }
func (ms *mockService2) Destroy() error   { return nil }
func (ms *mockService2) Get() interface{} { return mService{} }

func (ms *mockService3) Regist() error    { return nil }
func (ms *mockService3) Destroy() error   { return ErrTestDestroy }
func (ms *mockService3) Get() interface{} { return mService{} }

// main test to service registration
type testSuite struct {
	suite.Suite
	mockS *mockService1
}

func TestHDLSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (ts *testSuite) SetupSuite() {
	ts.mockS = &mockS1
}

func (ts *testSuite) SetupTest() {
	ServiceMap = make(map[string]*serviceCnf)

	Register(&mockS2)
}

func (ts *testSuite) TestRegiste() {
	testcase := []struct {
		Desc   string
		Svs    Service
		ExpErr error
	}{
		{
			Desc:   "Normal Case 1",
			Svs:    &mockS1,
			ExpErr: nil,
		},
		{
			Desc:   "Normal Case 2",
			Svs:    &mockS3,
			ExpErr: nil,
		},
		{
			Desc:   "service has been registe",
			Svs:    &mockS2,
			ExpErr: ErrServiceHasBeenRegisted,
		},
	}

	for _, c := range testcase {
		service := c.Svs
		err := Register(service)
		ts.Equal(c.ExpErr, err, c.Desc)
	}
}

func (ts *testSuite) TestGet() {
	testcase := []struct {
		Desc     string
		MockFunc func()
		Service  Service
		ExpRes   interface{}
		ExpErr   error
	}{
		{
			Desc:     "normal case",
			MockFunc: func() {},
			Service:  &mockS2,
			ExpRes:   mockS2.Get(),
			ExpErr:   nil,
		},
		{
			Desc:     "Not regist service, mockS2",
			MockFunc: func() {},
			Service:  &mockS3,
			ExpRes:   nil,
			ExpErr:   ErrServiceNotRegist,
		},
	}

	for _, c := range testcase {
		c.MockFunc()
		sv, err := Get(c.Service)
		ts.Equal(c.ExpErr, err, c.Desc)
		ts.Equal(c.ExpRes, sv, c.Desc)
	}
}

func (ts *testSuite) TestClean() {
	testcase := []struct {
		Desc     string
		MockFunc func()
		Name     string
		ExpErr   error
	}{
		{
			Desc:     "Clean not regist service",
			MockFunc: func() {},
			Name:     "what",
			ExpErr:   ErrServiceNotRegist,
		},
		{
			Desc: "clean mock Service 2",
			MockFunc: func() {
				Register(&mockS2)
			},
			Name:   "*mockService2",
			ExpErr: nil,
		},
		{
			Desc: "Clean service failure",
			MockFunc: func() {
				Register(&mockS3)
			},
			Name:   "*mockService3",
			ExpErr: ErrServiceDestroyFailure,
		},
	}

	for _, c := range testcase {
		c.MockFunc()
		err := Clean(c.Name)
		ts.Equal(c.ExpErr, err, c.Desc)
	}
}
