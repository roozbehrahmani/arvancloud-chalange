package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
)

const configFileDir = "config/config.json"

func Init(ctx context.Context) (*Config, error) {
	cnf := &Config{}
	configFile, err := os.Open(configFileDir)
	if err != nil {
		return cnf, err
	}

	defer configFile.Close()
	byteResult, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal([]byte(byteResult), cnf)
	if err != nil {
		return cnf, err
	}

	return cnf, nil
}
