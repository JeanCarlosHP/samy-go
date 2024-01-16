package viper

import (
	"fmt"

	"github.com/jeancarloshp/samy-go/pkg/config"
	"github.com/jeancarloshp/samy-go/pkg/database"
	"github.com/spf13/viper"
)

func LoadViperConfig() (*config.Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	return &config.Config{
		Enviroment: viper.GetString("ENV"),
		AppPort:    viper.GetString("APP_PORT"),
		Db: database.Db{
			Url:      viper.GetString("DB_URL"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
		},
	}, nil

}
