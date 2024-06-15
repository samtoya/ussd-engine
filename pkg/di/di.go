package di

import "sync"

type ServiceCollection interface {
	GetService(string) (interface{}, bool)
	Register(string, interface{}) error
}

type serviceCollection struct {
	sync.Mutex
	box map[string]interface{}
}

func (services *serviceCollection) Register(s string, i interface{}) error {
	services.Lock()
	defer services.Unlock()
	services.box[s] = i
	return nil
}

func (services *serviceCollection) GetService(s string) (interface{}, bool) {
	services.Lock()
	defer services.Unlock()
	value, exists := services.box[s]
	return value, exists
}

func NewServiceCollection() ServiceCollection {
	return &serviceCollection{
		box: make(map[string]interface{}),
	}
}
