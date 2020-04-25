package util

import (
	"encoding/csv"
	"log"
	"os"
)

const exportFldrName = "exports"

// Write is to write an array of mark 6 draw result to a target csv file
func Write(fileName string, data [][]string) {
	createExportFldr(exportFldrName)

	file, err := os.Create(exportFldrName + "/" + fileName)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func createExportFldr(fldrName string) {
	if _, err := os.Stat(fldrName); os.IsNotExist(err) {
		os.Mkdir(fldrName, os.ModeDir)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
