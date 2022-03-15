package gojsonconfig

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppId   int
	AppHash string
	Phone   string
	Min     int
}

func ReadConfig() Config {
	file, _ := os.Open("config.json")
	defer file.Close()
	config := Config{}
	error := json.NewDecoder(file).Decode(&config)
	if error != nil {
		panic(error)
	}
	return config
}
