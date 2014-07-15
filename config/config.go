package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Users  []string
	Groups []string
	Port   int
}

func AppConfiguration() (*Configuration, error) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Users)
	fmt.Println(configuration.Port, " port")
	return &configuration, nil
}
