package config

import (
	"time"

	"github.com/spf13/pflag"
)

type Config struct {
	ServerConfig
}

type ServerConfig struct {
	Host string
	Port int

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("APIConfig", pflag.PanicOnError)

	f.StringVar(&c.Host, "host", "127.0.0.1", "ip")
	f.IntVar(&c.Port, "port", 8081, "port")
	f.DurationVar(&c.ReadTimeout, "readtimeout", time.Duration(0), "api read timeout (default 0s)")
	f.DurationVar(&c.WriteTimeout, "writetimeout", time.Duration(0), "api write timeout (default 0s)")

	return f
}
