package entity

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

type FileInfo struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	CheckSum []byte `json:"check_sum"`
}

func GetFileInfo(filePath string) (FileInfo, error) {
	fileInfo := FileInfo{}
	fileInfo.FileName = filepath.Base(filePath)
	size, err := GetFileSize(filePath)
	if err != nil {
		return FileInfo{}, err
	}
	fileInfo.FileSize = size
	checkSum, err := CalculateCheckSum(filePath)
	if err != nil {
		return FileInfo{}, err
	}
	fileInfo.CheckSum = checkSum
	return fileInfo, nil
}

func GetFileSize(filePath string) (int64, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return -1, err
	}
	return fi.Size(), nil
}

func CalculateCheckSum(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Failed to calc checksum", err)
		}
	}(f)

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func (f *FileInfo) Marshal() ([]byte, error) {
	return json.Marshal(f)
}
