package usecase

import (
	"crypto/sha256"
	"file_transferring_server/internal/entity"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func CreateFilePath(filepath string, filename string, id int) string {
	return filepath + "/" + strconv.Itoa(id) + "_" + filename
}

func CreateFile(filepath string, filename string, id int) (*os.File, error) {
	file, err := os.Create(CreateFilePath(filepath, filename, id))
	if err != nil {
		return file, err
	}
	return file, nil
}

func CheckIfFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func CheckIfDirectoryExists(filepath string) bool {
	fileInfo, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir()
}

func CreateDirectory(filepath string) error {
	err := os.MkdirAll(filepath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func WriteToFile(file *os.File, dataPacket entity.DataPacket) error {
	_, err := file.Write(dataPacket.Data)
	if err != nil {
		return err
	}
	return nil
}

func GetFileInfo(filePath string) (entity.FileInfo, error) {
	fileInfo := entity.FileInfo{}
	fileInfo.FileName = filepath.Base(filePath)
	size, err := GetFileSize(filePath)
	if err != nil {
		return entity.FileInfo{}, err
	}
	fileInfo.FileSize = size
	checkSum, err := CalculateCheckSum(filePath)
	if err != nil {
		return entity.FileInfo{}, err
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

func CompareCheckSums(checkSum1 []byte, checkSum2 []byte) bool {
	if len(checkSum1) != len(checkSum2) {
		return false
	}
	for i := range checkSum1 {
		if checkSum1[i] != checkSum2[i] {
			return false
		}
	}
	return true
}
