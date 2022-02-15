package config

import (
	"encoding/json"
	"io/ioutil"
)

type UpbitKeyConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type Config struct {
	UpbitKey UpbitKeyConfig `json:"upbit_key"`
}

func LoadConfig(filePath string) (*Config, error) {
	cfg := &Config{}
	dataBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cfg, err
	}

	json.Unmarshal(dataBytes, cfg)

	return cfg, nil
}
