package apifilesystem

import (
	"github.com/sirupsen/logrus"
)

func StartFileSystem(config *ConfigDirectories) error {
	err := createDirictoriesAtTime(config)
	if err != nil {
		logrus.Fatal(err)
		return err
	}

	err = setCurrentContentDirPathAtTime(config)

	if err != nil {
		logrus.Fatal(config)
		return err
	}
	return nil
}
