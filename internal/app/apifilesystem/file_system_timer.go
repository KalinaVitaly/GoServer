package apifilesystem

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func setCurrentContentDirPathAtTime(config *ConfigDirectories) error {
	err := callAtTime(config, 0, 59, 59, 1, SetFilePathDirToSave)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func createDirictoriesAtTime(config *ConfigDirectories) error {
	err := callAtTime(config, 23, 59, 59, 24, CreateDirectoriesInFileSystem)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func callAtTime(config *ConfigDirectories, hour, min, sec int, timeToSleep int64, callerFunction func(config *ConfigDirectories) error) error {
	loc, err := time.LoadLocation("Local")

	if err != nil {
		return err
	}

	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * time.Duration(timeToSleep))
	}

	duration := firstCallTime.Sub(time.Now().Local())
	fmt.Println("Duration : ", duration)
	go func(config *ConfigDirectories) {
		time.Sleep(duration)
		for {
			if err := callerFunction(config); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Hour * time.Duration(timeToSleep))
		}
	}(config)

	return nil
}

func SetFilePathDirToSave(config *ConfigDirectories) error {
	currentDate := time.Now().Format("01-02-2006")
	currentTime := time.Now().Hour()
	config.CurrentDirectoryPath = config.ContentFilePath + currentDate + "_" + strconv.Itoa(currentTime) + "/"
	return nil
}
