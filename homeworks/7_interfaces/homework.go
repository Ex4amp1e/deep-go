package homework7

import "errors"

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	providers map[string]any
}

func NewContainer() *Container {
	return &Container{providers: make(map[string]any)}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.providers[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	service, ok := c.providers[name]
	if !ok {
		return nil, errors.New("not found")
	}
	constructor, ok := service.(func() any)
	if !ok {
		return nil, errors.New("wrong constructor function type")
	}
	return constructor(), nil
}
