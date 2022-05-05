package apifilesystem

type ConfigDirectories struct {
	ContentFilePath      string `toml:"dir_file_content"`
	DirectoriesCount     int    `toml:"directories_count"`
	CurrentDirectoryPath string
}

func NewConfigDirecories() *ConfigDirectories {
	return &ConfigDirectories{
		ContentFilePath:      "/",
		DirectoriesCount:     24,
		CurrentDirectoryPath: "",
	}
}
