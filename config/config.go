package config

import (
	"github.com/spf13/viper"
)

// LoadConfig loads the right config into the application based on the environment
func LoadConfig(env *string) error {
	viper.SetConfigName(*env)
	viper.AddConfigPath("./config/")

	return viper.ReadInConfig()
}

// GetString returns a string variable from the config file
func GetString(key string) string {
	return viper.GetString(key)
}
