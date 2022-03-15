package config

import (
	"encoding/json"
	"os"
)

func ReadConfig(filename string, targetStruct interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(targetStruct)
}
