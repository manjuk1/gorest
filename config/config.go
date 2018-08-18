package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Settings stores the application-wide configurations
var Settings appConfig

type appConfig struct {
	// JWT token secret
	Token          string `mapstructure:"tokenSecret"`
	
}

// LoadConfig loads configuration from the given list of paths and populates it into the Settings variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "GOREST" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("gorest")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Settings); err != nil {
		return err
	}
	return nil
}
