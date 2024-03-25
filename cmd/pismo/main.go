package main

import (
	"fmt"
	"self-projects/pismo/api/rest"
	"self-projects/pismo/configs"
)

func init() {
	configs.ReadConfig()
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
