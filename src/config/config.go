package config

import (
	"errors"
	"log"
	"os"
)

var FileName string

func LoadDBConfig() error {
	log.Println("INFO : Loading DB Config CSV File Name from env")
	FileName = os.Getenv("FileName")
	if FileName == "" {
		return errors.New("mandatory configuration missing for File Name from Config")
	}
	return nil
}
