package lib

import (
	"encoding/json"
	"os"
)

type CarConfigItem struct {
	CatalogTitle     string `json:"catalog_title"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	Name             string `json:"name"`
	Icon             string `json:"icon"`
	SubType          string `json:"sub_type"`
}

type CarConfig struct {
	CarName string          `json:"car_name"`
	Items   []CarConfigItem `json:"items"`
}

func LoadCarConfig(path string) (*CarConfig, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer configFile.Close()

	var carConfig CarConfig

	err = json.NewDecoder(configFile).Decode(&carConfig)

	if err != nil {
		return nil, err
	}

	return &carConfig, nil
}