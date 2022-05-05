package apifilesystem

//
//func CreateDirectoriesInFileSystem() error {
//	var dirName string
//	var dirPath string
//	currentData := time.Now().Format("01-02-2006")
//
//	for i := int64(0); i < DirCountToCreate+1; i++ {
//		dirName = currentData
//		dirPath = CurrentHomeFilePath + "/"
//		if i != 0 {
//			dirName = currentData + "_" + strconv.Itoa(int(i))
//			dirPath += currentData + "/"
//		}
//
//		errorDirCreate := createDirInFileSystem(dirPath, dirName)
//
//		if errorDirCreate != nil {
//			log.Fatal("Error: create dir failed!")
//			panic(errorDirCreate)
//		}
//	}
//
//	return nil
//}
