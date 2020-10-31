package di

/**
 * di package is a sample example which for dependency injection implementation in go practice
 */

import (
	"sync"
)

type serviceCnf struct {
	Service Service
	Once    *sync.Once
}

// Service describe module which want to set into dependency injection should implement function
type Service interface {
	// Regist to regist a service for dependency injection
	Regist() error

	// Destroy to remove registed service
	Destroy() error

	// Get to get service interface
	Get() interface{}
}
