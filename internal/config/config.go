package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Configurations struct {
	SdkConfig SDKConfigurations `mapstructure:"SDK_CONFIG"`
	DevConfig DevConfigurations `mapstructure:"DEV_CONFIG"`
}

type SDKConfigurations struct {
	SdkmanServiceUrl string `mapstructure:"SDKMAN_SERVICE"`
	SdkmanVersion    string `mapstructure:"SDKMAN_VERSION"`
	SdkmanDir        string `mapstructure:"SDKMAN_DIR"`
}

type DevConfigurations struct {
	Debug bool `mapstructure:"DEBUG"`
}

var (
	config *Configurations
	once   sync.Once
)

func GetConfigurations() *Configurations {
	once.Do(func() {
		viper.SetConfigName("sdk-config")
		viper.AddConfigPath(".\\config")
		viper.AutomaticEnv()
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error reading config file, %s", err)
		}

		err := viper.Unmarshal(&config)
		if err != nil {
			fmt.Println("Unable to decode into struct, %v", err)
		}
	})

	return config
}
