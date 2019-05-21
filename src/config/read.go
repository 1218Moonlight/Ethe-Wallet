package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

func File(path string) (param, error) {
	p := param{}

	jsonFile, err := os.Open(path)
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
