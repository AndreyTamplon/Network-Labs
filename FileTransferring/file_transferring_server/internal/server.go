package internal

import (
	"file_transferring_server/internal/entity"
	"file_transferring_server/internal/usecase"
	"file_transferring_server/pkg/connection"
	"file_transferring_server/pkg/logger"
	"file_transferring_server/pkg/tcp_server"
	"time"
)

const (
	_defaultNet = "tcp"
)

type Server struct {
	server *tcp_server.Server
	logger *logger.Logger
	port   string
}

func NewServer(port string, logger *logger.Logger) *Server {
	s := &Server{
		server: tcp_server.NewServer(),
		port:   port,
		logger: logger,
	}
	return s
}

func (server *Server) start() error {
	err := server.server.Listen(_defaultNet, server.port)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Run() error {
	if err := server.start(); err != nil {
		return err
	}
	server.logger.Info("Server started")
	defer func(s *tcp_server.Server) {
		err := s.Close()
		if err != nil {
			panic(err)
		}
	}(server.server)
	err := usecase.CreateDirectory("uploads")
	if err != nil {
		return err
	}
	for i := 0; ; i++ {
		conn, err := server.server.Accept()
		if err != nil {
			return err
		}
		server.logger.Info("New connection from %s", conn.RemoteAddr())
		go server.handleConnection(conn, i)
	}
}

// TODO split this function
func (server *Server) handleConnection(conn *connection.Connection, id int) {
	defer func() {
		err := conn.Close()
		if err != nil {
			server.logger.Warn("Error %s while closing connection %s", err, conn.RemoteAddr())
		}
	}()
	fileInfo := &entity.FileInfo{}
	data := make([]byte, 2048)
	n, err := conn.Read(data)
	if (err != nil) || (n == 0) {
		server.logger.Warn("Error %s while reading file info from %s", err, conn.RemoteAddr())
		return
	}
	err = fileInfo.Unmarshal(data[:n])
	if err != nil {
		server.logger.Warn("Error %s while unmarshalling file info from %s", err, conn.RemoteAddr())
		return
	}
	data = make([]byte, fileInfo.PacketLength)
	n = fileInfo.PacketLength
	server.logger.Info("File info received from %s", conn.RemoteAddr())
	file, err := usecase.CreateFile("uploads", fileInfo.FileName, id)
	if err != nil {
		server.logger.Warn("Error %s while creating file %s", err, fileInfo.FileName)
		return
	}
	err = conn.Write([]byte{1})
	if err != nil {
		server.logger.Warn("Error %s while sending confirmation to %s", err, conn.RemoteAddr())
		return
	}
	packet := &entity.DataPacket{}
	ticker := time.NewTicker(3 * time.Second)
	startTime := time.Now()
	receivedBytes := 0
	receivedBytesWithinLast3Seconds := 0
	for {
		n, err := conn.Read(data)
		if err != nil || n == 0 {
			server.logger.Warn("Error %s while reading file info from %s", err, conn.RemoteAddr())
			return
		}
		err = packet.Unmarshal(data[:n])
		if err != nil {
			server.logger.Warn("Error %s while unmarshalling packet from %s", err, conn.RemoteAddr())
			return
		}
		if packet.LastPacket {
			break
		}
		err = conn.Write([]byte{uint8(1)})
		if err != nil {
			server.logger.Warn("Error %s while sending confirmation to %s", err, conn.RemoteAddr())
			return
		}
		receivedBytes += len(packet.Data)
		receivedBytesWithinLast3Seconds += len(packet.Data)
		err = usecase.WriteToFile(file, *packet)
		if err != nil {
			server.logger.Warn("Error %s while writing to file %s", err, fileInfo.FileName)
			return
		}
		select {
		case <-ticker.C:
			server.logger.Info("Instantaneous velocity = %d Total velocity = %d\n",
				receivedBytesWithinLast3Seconds/3, receivedBytes/int(time.Since(startTime).Seconds()))
			receivedBytesWithinLast3Seconds = 0
		default:
			// do nothing
		}

	}
	server.logger.Info("File %s received from %s with total speed %d", fileInfo.FileName, conn.RemoteAddr(), receivedBytes/int(time.Since(startTime).Seconds()))
	err = file.Close()
	if err != nil {
		server.logger.Warn("Error %s while closing file %s", err, fileInfo.FileName)
		return
	}
	checkSum, err := usecase.CalculateCheckSum(usecase.CreateFilePath("uploads", fileInfo.FileName, id))
	if err != nil {
		server.logger.Warn("Error %s while calculating checksum for file %s", err, fileInfo.FileName)
		return
	}
	fileTransferredSuccessfully := uint8(1)
	if !usecase.CompareCheckSums(checkSum, fileInfo.CheckSum) {
		server.logger.Warn("Checksums don't match for file %s", fileInfo.FileName)
		fileTransferredSuccessfully = uint8(0)
	}
	server.logger.Info("Checksums match for file %s", fileInfo.FileName)
	err = conn.Write([]byte{fileTransferredSuccessfully})
	if err != nil {
		server.logger.Warn("Error %s while sending confirmation to %s", err, conn.RemoteAddr())
		return
	}
}
