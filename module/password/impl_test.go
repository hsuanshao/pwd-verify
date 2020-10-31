package password

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	im Service
}

func (t *testSuite) SetupSuite() {}

func (t *testSuite) SetupTest() {
}

func (t *testSuite) TearDownTest() {
	t.im = nil
}

func (t *testSuite) TestValidator() {
	type pwdTestCase struct {
		Desc         string
		Password     string
		ExpLength    bool
		ExpUppercase bool
		ExpLowercase bool
		ExpNumber    bool
		ExpSymbol    bool
		ExpSequence  bool
		ExpErr       error
	}

	testCase := []struct {
		Desc  string
		Test  func()
		cases []pwdTestCase
	}{
		{
			Desc: "5-12 length, lowercase and number allowed",
			Test: func() {
				t.im = New(5, 12, false, true, true, false, false)
			},
			cases: []pwdTestCase{
				{
					Desc:         "correct password format",
					Password:     "qwb3sd15",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       nil,
				},
				{
					Desc:         "has sequence character",
					Password:     "qab3sd15",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  false,
					ExpErr:       ErrFormat,
				},
				{
					Desc:         "has revert sequence character",
					Password:     "q321sd15",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  false,
					ExpErr:       ErrFormat,
				},
				{
					Desc:         "require lowercase character",
					Password:     "25282596",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: false,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       ErrFormat,
				},
				{
					Desc:         "incorrect string length",
					Password:     "2c",
					ExpLength:    false,
					ExpUppercase: true,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       ErrFormat,
				},
				{
					Desc:         "unexpect uppercase",
					Password:     "qwb3Dd1",
					ExpLength:    true,
					ExpUppercase: false,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       ErrFormat,
				},
			},
		},
		{
			Desc: "5-8 length, lowercase, uppercase, number, sequence, symbol allowed",
			Test: func() {
				t.im = New(5, 8, true, true, true, true, true)
			},
			cases: []pwdTestCase{
				{
					Desc:         "correct password format",
					Password:     "q-Wbd15",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       nil,
				},
				{
					Desc:         "require lowercase character",
					Password:     "22-A865",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: false,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       ErrFormat,
				},
				{
					Desc:         "incorrect string length",
					Password:     "2c-Bsdadfds",
					ExpLength:    false,
					ExpUppercase: true,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       ErrFormat,
				},
			},
		},
		{
			Desc: "5-12 length, uppercase, number, sequence, symbol allowed",
			Test: func() {
				t.im = New(5, 12, true, false, true, true, true)
			},
			cases: []pwdTestCase{
				{
					Desc:         "correct password format",
					Password:     "H-WBD(15)",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: true,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       nil,
				},
				{
					Desc:         "lowercase character not allowd",
					Password:     "22b-A865",
					ExpLength:    true,
					ExpUppercase: true,
					ExpLowercase: false,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       ErrFormat,
				},
				{
					Desc:         "incorrect string length and has lowercase character",
					Password:     "2c-Bsdasdbfsd",
					ExpLength:    false,
					ExpUppercase: true,
					ExpLowercase: false,
					ExpNumber:    true,
					ExpSymbol:    true,
					ExpSequence:  true,
					ExpErr:       ErrFormat,
				},
			},
		},
	}

	for _, c := range testCase {
		c.Test()
		for _, pwdCase := range c.cases {
			lv, ucv, lcv, nv, symbolv, sv, err := t.im.Validator(pwdCase.Password)

			t.Equal(pwdCase.ExpLength, lv, "expect password length rule", c.Desc, pwdCase.Desc)
			t.Equal(pwdCase.ExpUppercase, ucv, "expect uppwercase rule", c.Desc, pwdCase.Desc)
			t.Equal(pwdCase.ExpLowercase, lcv, "expect lowercase rule", c.Desc, pwdCase.Desc)
			t.Equal(pwdCase.ExpNumber, nv, "expect number rule", c.Desc, pwdCase.Desc)
			t.Equal(pwdCase.ExpSymbol, symbolv, "expect symbol rule", c.Desc, pwdCase.Desc)
			t.Equal(pwdCase.ExpSequence, sv, "expect sequence rule", c.Desc, pwdCase.Desc)
			t.Equal(pwdCase.ExpErr, err, "expect error", c.Desc, pwdCase.Desc)
		}
		t.TearDownTest()
	}
}

func TestService(t *testing.T) {
	suite.Run(t, new(testSuite))
}
