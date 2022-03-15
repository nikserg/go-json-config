package config

import (
	"testing"
)

type Config1 struct {
	X int
	Y int
}

type Config2 struct {
	X string
	Y int
}

func Test1(t *testing.T) {
	var config Config1
	err := ReadConfig("test/config1.json", &config)
	if err != nil {
		t.Error(err)
	}
	if config.X != 123 {
		t.Error()
	}
	if config.Y != 456 {
		t.Error()
	}
}
func Test2(t *testing.T) {
	var config Config2
	err := ReadConfig("test/config2.json", &config)
	if err != nil {
		t.Error(err)
	}
	if config.X != "abc" {
		t.Error()
	}
	if config.Y != 456 {
		t.Error()
	}
}