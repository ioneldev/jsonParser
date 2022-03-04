package service

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadObjectFromJson(bytes []byte, object interface{}) error {
	err := json.Unmarshal(bytes, object)

	if err != nil {
		fmt.Println("failed to decode: " + err.Error())
		return err
	}

	return nil
}

func EncodeJson(object interface{}) ([]byte, error) {
	data, err := json.Marshal(object)

	if err != nil {
		fmt.Println("failed to decode: " + err.Error())
		return nil, err
	}

	return data, nil
}

func WriteJsonFile(data []byte, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
