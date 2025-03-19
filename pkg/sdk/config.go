package sdk

import "time"

// DefaultConfig is the default configuration to use.
var DefaultConfig = &Config{
	ExpireDelay: 5 * time.Minute,
	GrpcAddr:    "localhost:8081",
}

// Config represents the SDK configuration values.
type Config struct {
	ExpireDelay time.Duration
	GrpcAddr    string

	ServiceName    string
	RegistrySchema string
	Endpoints      []string

	ClientID     string
	ClientSecret string
}
