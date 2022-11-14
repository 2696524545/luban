package utils

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/dnsjia/luban/cmd/options"
)

const (
	defaultConfigFile = "etc/config.yaml"
)

func Viper(path string) *viper.Viper {
	config := viper.New()
	if path != "" {
		config.SetConfigFile(path)
	} else {
		config.SetConfigFile(defaultConfigFile)
	}

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal err config file: %s\n", err))
	}

	if err := config.Unmarshal(&options.Config); err != nil {
		fmt.Println(err)
	}

	return config
}
