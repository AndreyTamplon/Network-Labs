package app

import (
	"file_transferring_client/config"
	"file_transferring_client/internal/usecase"
	"file_transferring_client/pkg/connection"
	"file_transferring_client/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	connection := connection.NewConnection()
	if err := connection.Connect("tcp", cfg.Network.IpAddress+":"+cfg.Network.Port); err != nil {
		l.Fatal(err)
	}
	defer func() {
		if err := connection.Close(); err != nil {
			l.Fatal(err)
		}
	}()

	err := usecase.SendFileInfo(connection, cfg.Data.FilePath, l)
	if err != nil {
		l.Fatal(err)
	}
	delivery, err := usecase.GetFileDeliveryInformation(connection)
	if err != nil {
		l.Fatal(err)
	}
	err = usecase.SendFile(connection, cfg.Data.FilePath, l)
	if err != nil {
		l.Fatal(err)
	}
	delivery, err = usecase.GetFileDeliveryInformation(connection)
	if err != nil {
		l.Fatal(err)
	}
	if delivery {
		l.Info("File was successfully delivered")
	} else {
		l.Info("File was not delivered")
	}

}
