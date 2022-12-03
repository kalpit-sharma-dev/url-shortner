package db

import (
	"log"
	"os"

	"github.com/kalpit-sharma-dev/url-shortner/src/config"
)

func GetFileProvider() (csvFile *os.File) {

	csvFile, err := os.OpenFile("../../"+config.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error creating csv file", err)
	}
	//defer csvFile.Close()
	return csvFile
}
