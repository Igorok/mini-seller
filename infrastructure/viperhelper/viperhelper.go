// Package viperhelper - helper for viper, uses to apply configuration from local files and environment
package viperhelper

import (
	"os"

	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

// IViper - helper for viper, describe methods for initiation of config
type IViper interface {
	updateSettings()
	Read() error
	getEnv()
}

// default values for json config
const (
	configType = "json"
	configName = "config"
	configPath = "./infrastructure/viperhelper"
)

// list of environment variables
var envVariables = []string{
	"MONGO_DB", "MONGO_HOST", "MONGO_PORT", "MONGO_USER",
	"MONGO_PASSWORD", "MONGO_AUTH", "MONGO_REPLICASET",
	"WEB_PORT",
}

// Viper - class for initialization of viper config with values from local configuration and environment
type Viper struct {
	ConfigType, ConfigName, ConfigPath string
}

// updateSettings - set default values for arguments
func (vip *Viper) updateSettings() {
	if vip.ConfigType == "" {
		vip.ConfigType = configType
	}
	if vip.ConfigName == "" {
		vip.ConfigName = configName
	}
	if vip.ConfigPath == "" {
		vip.ConfigPath = configPath
	}
}

// getEnv - get variables from environment
func (vip *Viper) getEnv() {
	for _, variable := range envVariables {
		value := os.Getenv(variable)
		if value != "" {
			viper.Set(variable, value)
		}
	}
}

// Read - read configuration
func (vip *Viper) Read() error {
	vip.updateSettings()

	viper.SetConfigType(vip.ConfigType)
	viper.AddConfigPath(vip.ConfigPath)

	viper.SetConfigName("local-" + vip.ConfigName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Info("Local config", err)
		viper.SetConfigName(vip.ConfigName)
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	vip.getEnv()

	return nil
}
