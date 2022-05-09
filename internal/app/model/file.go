package model

type File struct {
	ID            int    `json:"id"`
	FileOwner     int    `json:"file_owner"`
	FilePath      string `json:"file_path"`
	FileName      string `json:"file_name"`
	FileQuery     string `json:"file-query"`
	FileAvailable bool   `json:"file_available"`
}
