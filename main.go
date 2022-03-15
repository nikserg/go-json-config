package config

import (
	"encoding/json"
	"os"
)

func ReadConfig(filename string, targetStruct interface{}) error {
	file, _ := os.Open(filename)
	defer file.Close()
	return json.NewDecoder(file).Decode(targetStruct)
}
