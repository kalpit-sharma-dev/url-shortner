package db

import (
	"log"
	"os"
)

func GetFileProvider() (csvFile *os.File) {

	csvFile, err := os.OpenFile("../../rrt.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error creating csv file", err)
	}
	//defer csvFile.Close()
	return csvFile
}
