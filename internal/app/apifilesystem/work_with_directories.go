package apifilesystem

import (
	"Diplom/internal/app/store"
	"io/ioutil"
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

		errorDirCreate := CreateDirInFileSystem(dirPath, dirName)

		if errorDirCreate != nil {
			log.Fatal("Error: create dir failed!")

		}
	}

	return nil
}

func CreateDirInFileSystem(filePath, dirName string) error {
	err := os.Mkdir(filePath+dirName, 444)
	if err != nil {
		return err
	}

	return nil
}

func SaveFileInDirectory(data []byte) error {
	file, err := os.Create(NewConfigDirecories().ContentFilePath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	writeFileLength, isOk := file.Write(data)

	if isOk != nil || writeFileLength != len(data) {
		log.Fatal("Error: write file failed!")
		return store.ErrWriteFile
	}
	return nil
}

func ReadFile(fileQuery string) ([]byte, error) {
	if !IsFileExists(fileQuery) {
		return nil, store.ErrFoundFile
	}

	data, err := ioutil.ReadFile(fileQuery)
	if err != nil {
		log.Fatal(err)
		return nil, store.ErrReadFile
	}

	return data, nil
}

func DeleteFile(filePath string) error {
	if isOk := IsFileExists(filePath); !isOk {
		return store.ErrFoundFile
	}

	if err := os.Remove(filePath); err != nil {
		return store.ErrFoundFile
	}

	return nil
}

func IsFileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) || fileInfo.IsDir() {
		return false
	}
	return true
}
