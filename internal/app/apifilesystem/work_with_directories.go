package apifilesystem

import (
	"log"
	"os"
	"strconv"
	"time"
)

func CreateDirectoriesInFileSystem(config *ConfigDirectories) error {
	var dirName string
	var dirPath string
	currentData := time.Now().Format("01-02-2006")

	for i := 0; i < config.DirectoriesCount+1; i++ {
		dirName = currentData
		dirPath = config.ContentFilePath
		if i != 0 {
			dirName = currentData + "_" + strconv.Itoa(int(i))
			dirPath += currentData + "/"
		}

		errorDirCreate := createDirInFileSystem(dirPath, dirName)

		if errorDirCreate != nil {
			log.Fatal("Error: create dir failed!")

		}
	}

	return nil
}

func createDirInFileSystem(filePath, dirName string) error {
	err := os.Mkdir(filePath+dirName, 444)
	if err != nil {
		log.Fatal("Create dir error")
		return err
	}

	return err
}
