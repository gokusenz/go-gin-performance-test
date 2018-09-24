package config

import (
	"encoding/json"
	"io/ioutil"

	"fastwork/go-gin-performance-test/app/models"
)

// Config variable
var Config models.Config

func init() {
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &Config); err != nil {
		panic(err)
	}
}
