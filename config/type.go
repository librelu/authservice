package config

import "github.com/spf13/viper"

// Handler Config handler
type Handler interface {
	GetString(key, fallback string) string
}

type config struct {
	config *viper.Viper
}
