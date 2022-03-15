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
type Config3 struct {
	X string `json:"x"`
	Y int    `json:"y"`
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
func Test3(t *testing.T) {
	var config Config3
	err := ReadConfig("test/config3.json", &config)
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
func Test4(t *testing.T) {
	var config Config3
	err := ReadConfig("kek.json", &config)
	if err == nil {
		t.Error("No error on non-existing file")
	}
}
