package service

import (
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	Ready          bool
	ApiUrl         string `json:"ApiUrl"`
	ApiLinesToRead int    `json:"ApiLinesToRead"`
}

func (conf *Config) Initialize(configFile string) error {
	jsonFile, err := os.Open(configFile)

	if err != nil {
		return errors.New("failed to open config file")
	}

	bytes, _ := ioutil.ReadAll(jsonFile)

	err = LoadObjectFromJson(bytes, conf)

	if err != nil {
		return errors.New("failed to decode config json")
	}

	err = jsonFile.Close()

	if err != nil {
		return errors.New("failed to close config file")
	}

	conf.Ready = true
	return nil
}
