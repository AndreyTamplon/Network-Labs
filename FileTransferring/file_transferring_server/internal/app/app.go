package app

import (
	"file_transferring_server/config"
	"file_transferring_server/internal"
	"file_transferring_server/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	server := internal.NewServer(cfg.Port, l)
	if err := server.Run(); err != nil {
		l.Fatal(err)
	}
}
