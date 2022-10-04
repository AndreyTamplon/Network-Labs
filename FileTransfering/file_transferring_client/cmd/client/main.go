package main

import (
	"file_transferring_client/config"
	"file_transferring_client/internal/app"
	"log"
	"net"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("usage: <host> <port> <file>\n")
	}
	ipAddress := os.Args[1]
	if net.ParseIP(ipAddress) == nil {
		log.Fatalf("invalid IP address: %s\n", ipAddress)
	}
	port := os.Args[2]
	if match, err := regexp.Match("^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))$", []byte(port)); err != nil || !match {
		log.Fatalf("invalid port: %s\n", port)
	}

	if _, err := os.Stat(os.Args[3]); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("file not found: %s\n", os.Args[1])
		} else {
			log.Fatalf("error: %s\n", err)
		}
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	cfg.Network.IpAddress = ipAddress
	cfg.Network.Port = port
	cfg.Data.FilePath = os.Args[3]
	app.Run(cfg)
}
