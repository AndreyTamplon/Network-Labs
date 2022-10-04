package entity

import (
	"encoding/json"
)

type FileInfo struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	CheckSum []byte `json:"check_sum"`
}

func (f *FileInfo) Marshal() ([]byte, error) {
	return json.Marshal(f)
}

func (f *FileInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, f)
}
