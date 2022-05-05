package apifilesystem

var configDir *ConfigDirectories

type ConfigDirectories struct {
	ContentFilePath      string `toml:"dir_file_content"`
	DirectoriesCount     int    `toml:"directories_count"`
	CurrentDirectoryPath string
}

func NewConfigDirecories() *ConfigDirectories {
	if configDir != nil {
		return configDir
	}

	configDir = &ConfigDirectories{
		ContentFilePath:      "/",
		DirectoriesCount:     24,
		CurrentDirectoryPath: "",
	}

	return configDir
}
