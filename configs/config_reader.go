package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func ReadConfig() {
	configName := os.Getenv("HOST_TYPE")
	if len(configName) == 0 {
		configName = "local"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error while reading the config")
	}
	LoadConfigFromEnvVariables()
}

func LoadConfigFromEnvVariables() {
	for _, key := range viper.AllKeys() {
		configValue := viper.Get(key)
		viperKey := strings.ToUpper(key)
		formattedKey := strings.ReplaceAll(viperKey, ".", "_")
		if value, present := os.LookupEnv(formattedKey); present {
			configValue = value
		}

		viper.Set(key, configValue)
	}
}
