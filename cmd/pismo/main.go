package main

import (
	"fmt"
	"self-projects/pismo/api/rest"
	"self-projects/pismo/configs"
	"self-projects/pismo/pkg/logger"

	"github.com/spf13/viper"
)

func init() {
	configs.ReadConfig()
	logger.InitLogger(viper.GetString(configs.LoggingFormat), viper.GetString(configs.LoggingLevel))
}

func main() {
	fmt.Println("starting a server")
	s := rest.BuildServer()
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("error while starting the server")
		return
	}
	fmt.Println("closing the server")
}
