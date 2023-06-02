package utils

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() (config Config, err error) {
	path, err := os.Getwd()

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}
