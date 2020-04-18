package config

import (
	"github.com/spf13/viper"
)

// NewConfig getting config file
func NewConfig() (h Handler, err error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("$GOPATH/src/github.com/authsvc/config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return &config{
		config: v,
	}, nil
}

func (c config) GetString(key, fallback string) string {
	return c.config.GetString(key)
}
