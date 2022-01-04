// Package config has a configuration structure
package config

// Config contains configuration data
type Config struct {
	HostGrpc string `env:"HOST_GRPC" envDefault:"localhost"`
	PortGrpc string `env:"PORT_GRPC" envDefault:"10000"`
}
