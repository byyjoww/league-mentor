package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

func Build() Config {
	vpr, err := ConfigureViper("config/config.yaml", "app", "yaml")
	if err != nil {
		panic("failed to create viper instance")
	}

	return parse(vpr)
}

func parse(vpr *viper.Viper) Config {
	value := *new(Config)
	if err := vpr.Unmarshal(&value); err != nil {
		panic("failed to unmarshal full config: " + err.Error())
	}
	return value
}

func ConfigureViper(path, envPrefix, configType string) (*viper.Viper, error) {
	vpr := viper.New()
	defaultConfig := bytes.NewReader([]byte{})
	vpr.SetConfigType("yaml")
	if err := vpr.MergeConfig(defaultConfig); err != nil {
		return nil, err
	}

	// Override config
	vpr.SetConfigFile(path)
	vpr.SetConfigType(configType)
	if err := vpr.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return nil, err
		}
		// Dont return error if file is missing as overwrite is optional
	}

	// Override env variables
	vpr.AutomaticEnv()
	vpr.SetEnvPrefix(envPrefix)
	vpr.AddConfigPath(".")
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Manually set all config values, only needed when using viper.Unmarshal()
	for _, key := range vpr.AllKeys() {
		val := vpr.Get(key)
		vpr.Set(key, val)
	}

	return vpr, nil
}
