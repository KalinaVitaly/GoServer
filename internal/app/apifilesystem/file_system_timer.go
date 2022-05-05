package apifilesystem

import (
	"fmt"
	"time"
)

func CallAtTimeToCreateDir(homeDirPath string, count, hour, min, sec int, timeToSleep int64, callerFunction func() error) error {
	loc, err := time.LoadLocation("Local")

	if err != nil {
		return err
	}

	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), 23, 59, 30, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * time.Duration(timeToSleep))
	}

	duration := firstCallTime.Sub(time.Now().Local())
	fmt.Println("Duration : ", duration)
	go func() {
		time.Sleep(duration)
		for {
			callerFunction()
			time.Sleep(time.Hour * time.Duration(timeToSleep))
		}
	}()

	return nil
}

//func SetFilePathDirToSave() error {
//	currentDate := time.Now().Format("01-02-2006")
//	currentTime := time.Now().Hour()
//	CurrentFilePath = CurrentHomeFilePath + currentDate + "_" + strconv.Itoa(currentTime) + "/"
//	return nil
//}
