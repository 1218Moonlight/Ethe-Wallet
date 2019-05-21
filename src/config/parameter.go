package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

const jsonPath string = "config.json"

type param struct {
	Title       string `json:"title"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	HashMenubar bool   `json:"hasMenubar"`
}

func NewConfig() (param, error) {
	p := param{}

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return p, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return p, err
	}

	json.Unmarshal(byteValue, &p)

	return p, nil
}
