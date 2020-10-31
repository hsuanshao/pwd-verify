package di

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	createTimeout = time.Minute
)

var (
	// ServiceMap stores all service singletion
	ServiceMap = make(map[string]*serviceCnf)

	// ErrServiceHasBeenRegisted - describe service is exists
	ErrServiceHasBeenRegisted = fmt.Errorf("Service has been registed")
	// ErrServiceNotRegist - service does not regist
	ErrServiceNotRegist = fmt.Errorf("Service not regist")
	// ErrServiceDestroyFailure - destroy service failure
	ErrServiceDestroyFailure = fmt.Errorf("Service destroy failure")
	// ErrRegisterNilService describe caller request a nil service
	ErrRegisterNilService = fmt.Errorf("Can not register a nil service")
)

// Register handle service registration
func Register(sv Service) error {
	if sv == nil {
		panic(ErrRegisterNilService)
	}

	name := getName(sv)
	if _, ok := ServiceMap[name]; ok {
		logrus.WithFields(logrus.Fields{
			"service": name,
		}).Error(ErrServiceHasBeenRegisted)
		return ErrServiceHasBeenRegisted
	}

	ServiceMap[name] = &serviceCnf{
		Service: sv,
		Once:    &sync.Once{},
	}
	return nil
}

// Get handle get service interface
func Get(sv Service) (interface{}, error) {
	name := getName(sv)

	if _, ok := ServiceMap[name]; !ok {
		return nil, ErrServiceNotRegist
	}

	// createService will only be executed once
	ServiceMap[name].Once.Do(func() {
		createService(name)
	})

	return ServiceMap[name].Service.Get(), nil
}

// Clean to destory service by given service name
func Clean(serviceName string) error {
	if _, ok := ServiceMap[serviceName]; !ok {
		return ErrServiceNotRegist
	}

	srv := ServiceMap[serviceName]
	err := srv.Service.Destroy()
	if err != nil {
		return ErrServiceDestroyFailure
	}
	ServiceMap[serviceName].Once = &sync.Once{}
	return nil
}

func createService(name string) {
	start := time.Now()
	logrus.WithField("service", name).Info("creating service")
	done := make(chan struct{}, 1)
	go func() {
		ServiceMap[name].Service.Regist()
		done <- struct{}{}
	}()
	select {
	case <-done:
		logrus.WithFields(logrus.Fields{
			"service":  name,
			"duration": time.Since(start).Milliseconds(),
		}).Info("service created")
	case <-time.After(createTimeout):
		panic(fmt.Sprintf("create service %s failed, time out %s", name, createTimeout))
	}

}

func getName(sv Service) string {
	t := reflect.TypeOf(sv)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}
