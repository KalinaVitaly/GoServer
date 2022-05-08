package apifilesystem

import (
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
			filePath: "test.txt",
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
