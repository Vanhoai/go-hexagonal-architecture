package di

import (
	"fmt"
	"reflect"
	"sync"
)

// Container is a dependency injection container interface.
type DIContainer struct {
	mutex     sync.RWMutex
	services  map[string]interface{}
	providers map[string]reflect.Value
}

// CreateContainer creates a new instance of Container.
func CreateContainer() *DIContainer {
	return &DIContainer{
		services:  make(map[string]interface{}),
		providers: make(map[string]reflect.Value),
	}
}

// Register registers a service with the factory function.
func (c *DIContainer) Register(name string, provider interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	providerValue := reflect.ValueOf(provider)
	if providerValue.Kind() != reflect.Func {
		return fmt.Errorf("provider must be a function")
	}

	c.providers[name] = providerValue
	return nil
}

// Singleton registers a singleton service with the factory function.
func (c *DIContainer) Singleton(name string, provider interface{}) error {
	return c.Register(name, provider)
}

// Get retrieves a service by its name.
func (c *DIContainer) Get(name string) (interface{}, error) {
	c.mutex.RLock()

	// Check if the service is already instantiated
	if service, exists := c.services[name]; exists {
		c.mutex.RUnlock()
		return service, nil
	}

	provider, exists := c.providers[name]
	c.mutex.RUnlock()
	if !exists {
		return nil, fmt.Errorf("service not found: %s", name)
	}

	// Create the service instance
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Double-check if the service was instantiated while acquiring the write lock
	if service, exists := c.services[name]; exists {
		return service, nil
	}

	// Check if the provider exists
	results := provider.Call([]reflect.Value{})
	if len(results) == 0 {
		return nil, fmt.Errorf("provider function must return a value")
	}

	service := results[0].Interface()
	c.services[name] = service
	return service, nil
}

// Resolve automatically resolves and injects dependencies into the provided function.
func (c *DIContainer) Resolve(target interface{}) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		return fmt.Errorf("target must be a pointer")
	}

	targetValue = targetValue.Elem()
	if targetValue.Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to struct")
	}

	targetType := targetValue.Type()
	for i := 0; i < targetValue.NumField(); i++ {
		field := targetValue.Field(i)
		fieldType := targetType.Field(i)

		injectTag := fieldType.Tag.Get("inject")
		if injectTag == "" {
			continue
		}

		service, err := c.Get(injectTag)
		if err != nil {
			return fmt.Errorf("failed to inject %s: %w", injectTag, err)
		}

		if field.CanSet() {
			serviceValue := reflect.ValueOf(service)
			if serviceValue.Type().AssignableTo(field.Type()) {
				field.Set(serviceValue)
			}
		}
	}

	return nil
}
