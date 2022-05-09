package apifilesystem

import (
	"Diplom/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsFileExists(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		wanted   bool
	}{
		{
			name:     "check file text",
			filePath: "C:\\Users\\Kalina\\Desktop\\book\\configFileSystem.txt",
			wanted:   true,
		},
		{
			name:     "check directory",
			filePath: "internal/app/apifilesystem",
			wanted:   false,
		},
		{
			name:     "check unknown file",
			filePath: "internal/app/apifilesystem/filesystem_test_with_unknown_way.go",
			wanted:   false,
		},
		{
			name:     "model file",
			filePath: "C:\\Users\\Kalina\\Desktop\\book\\Effektivnoe-programmirovanie-TCP-IP_RuLit_Me_606683.pdf",
			wanted:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wanted, IsFileExists(tc.filePath))
		})
	}
}

func TestCreateDirectoriesInFileSystem(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		fileName string
		wanted   error
	}{
		{
			name:     "create file in content folder",
			filePath: "C:\\Users\\Kalina\\Desktop\\Diplom\\content\\",
			fileName: "data_new_content",
			wanted:   nil,
		},
		{
			name:     "create mp4 file in content folder",
			filePath: "C:\\Users\\Kalina\\Desktop\\Diplom\\content\\",
			fileName: "data.mp4",
			wanted:   nil,
		},
		{
			name:     "create png file in content folder",
			filePath: "C:\\Users\\Kalina\\Desktop\\Diplom\\content\\",
			fileName: "data_folder",
			wanted:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			assert.Equal(t, tc.wanted, CreateDirInFileSystem(tc.filePath, tc.fileName))
		})
	}
}

func TestReadFile(t *testing.T) {
	testCases := []struct {
		name          string
		filePath      string
		wanted        error
		fileSize      int
		containsError bool
	}{
		{
			name:          "read pdf file effectivnoe tcp ip programmirovanie",
			filePath:      "C:\\Users\\Kalina\\Desktop\\book\\Effektivnoe-programmirovanie-TCP-IP_RuLit_Me_606683.pdf",
			wanted:        nil,
			fileSize:      1526719,
			containsError: false,
		},
		{
			name:          "read pdf file Komandnaya stoka linux",
			filePath:      "C:\\Users\\Kalina\\Desktop\\book\\Komandnaya_stroka_Linux.pdf",
			wanted:        nil,
			fileSize:      5456716,
			containsError: false,
		},
		{
			name:          "create png file in content folder",
			filePath:      "C:\\Users\\Kalina\\Desktop\\Diplom\\content\\Drawing1_2025.bak",
			wanted:        store.ErrFoundFile,
			fileSize:      0,
			containsError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.containsError {
				assert.Equal(t, tc.wanted, func(filePath string) error {
					_, err := ReadFile(filePath)
					if err != nil {
						return err
					}
					return nil
				}(tc.filePath))
			} else {
				assert.Equal(t, tc.fileSize, func(filePath string) int {
					data, err := ReadFile(filePath)
					if err != nil {
						return -1
					}
					return len(data)
				}(tc.filePath))
			}
		})
	}
}
