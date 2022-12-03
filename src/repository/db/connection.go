package db

import (
	"log"
	"os"
)

func GetFileProvider() (csvFile *os.File) {

	//csvFile, err := os.Create("database.csv")
	csvFile, err := os.OpenFile("../../database.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		log.Fatal("error creating csv file")
	}
	//defer csvFile.Close()
	return csvFile
}
