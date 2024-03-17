package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func Read(config any) error {
	fileName := os.Getenv("CONFIG_FILE")
	if fileName == "" {
		fileName = "config.json"
	}
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	return ReadFile(config, fileName)
}

func ReadFile(config any, fileName string) error {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	buf, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, config)
	if err != nil {
		return err
	}
	log.Printf("Loaded %s config: %+v", fileName, config)
	return nil
}
