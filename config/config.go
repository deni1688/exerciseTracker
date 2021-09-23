package config

import (
	"flag"
	"github.com/spf13/viper"
)

// LoadConfigFile loads the right config into the application based on the environment
func LoadConfigFile(env string) error {
	viper.SetConfigName(env)
	viper.AddConfigPath("./config/")

	return viper.ReadInConfig()
}

// GetString returns a string variable from the config file
func GetString(key string) string {
	return viper.GetString(key)
}

func GetFlags() string {
	env := flag.String("env", "testing", "sets the config environment")
	flag.Parse()
	return *env
}
