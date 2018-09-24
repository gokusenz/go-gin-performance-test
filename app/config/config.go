package config

import (
	"fastwork/go-gin-performance-test/app/models"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config model
var Config models.Config

func init() {
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(file, &Config); err != nil {
		panic(err)
	}
}

// GetConfig is a function to retrive config
func GetConfig() *models.Config {
	return &Config
}
