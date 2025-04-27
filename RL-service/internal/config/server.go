package config

type PathConfig struct {
	Port string `envconfig:"PORT" default:"8080"`
}
