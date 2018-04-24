package basehttp

import "fmt"

// Service provides service related functionality
type Service struct {
	config Config
}

// NewService creates a new Service
func NewService(c Config) Service {
	return Service{
		config: c,
	}
}

// Greet will return a greeting
func (s Service) Greet(name string) (string, error) {
	return fmt.Sprintf("Oh, hello %s", name), nil
}
