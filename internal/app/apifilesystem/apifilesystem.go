package apifilesystem

func StartFileSystem(config *ConfigDirectories) error {
	err := createDirictoriesAtTime(config)
	if err != nil {
		return err
	}

	err = setCurrentContentDirPathAtTime(config)

	if err != nil {
		return err
	}
	return nil
}
