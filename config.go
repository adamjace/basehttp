package basehttp

import (
	"os"
)

const (
	defaultBind = ":9876"
)

// Config holds the configuration for the service
type Config struct {
	Bind string
}

// NewConfig returns a new Config struct
func NewConfig() Config {
	config := Config{
		Bind: os.Getenv("BIND"),
	}

	if config.Bind == "" {
		config.Bind = defaultBind
	}

	return config
}
