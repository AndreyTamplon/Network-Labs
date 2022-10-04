package usecase

import (
	"file_transferring_client/internal/entity"
	"file_transferring_client/pkg/connection"
	"file_transferring_client/pkg/logger"
	"io"
	"log"
	"os"
)

var packageSize = 1024

func SendFile(connection connection.Connection, filePath string, logger *logger.Logger) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Failed to close file", err)
		}
	}(file)
	data := make([]byte, packageSize)
	packetDelivery := make([]byte, 1)
	for {
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				logger.Info("File fully read")
				break
			}
			logger.Error("Failed to read file", err)
			return err
		}
		packet := entity.NewPacket(data[:n])
		marshalPacket, err := packet.Marshal()
		if err != nil {
			logger.Error("Failed to marshal packet", err)
			return err
		}
		err = connection.Write(marshalPacket)
		if err != nil {
			logger.Error("Failed to send packet", err)
			return err
		}
		n, err = connection.Read(packetDelivery)
		if err != nil || n != 1 {
			logger.Error("Failed to receive packet delivery confirmation", err)
			return err
		}
		if packetDelivery[0] != 1 {
			logger.Error("Packet was not delivered")
			return err
		}

	}
	//send packet with empty data and last packet flag
	packet := entity.NewPacket([]byte{})
	packet.LastPacket = true
	marshalPacket, err := packet.Marshal()
	if err != nil {
		return err
	}
	err = connection.Write(marshalPacket)
	if err != nil {
		return err
	}
	return nil
}

func SendFileInfo(connection connection.Connection, filePath string, logger *logger.Logger) error {
	fileInfo, err := entity.GetFileInfo(filePath)
	if err != nil {
		return err
	}
	data, err := fileInfo.Marshal()
	if err != nil {
		logger.Error("Failed to marshal file info", err)
		return err
	}
	return connection.Write(data)
}

func GetFileDeliveryInformation(connection connection.Connection) (bool, error) {
	data := make([]byte, 1)
	n, err := connection.Read(data)
	if err != nil || n != 1 {
		return false, err
	}
	return data[0] == 1, nil
}
