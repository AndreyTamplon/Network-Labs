package tcp_server

import (
	"file_transferring_server/pkg/connection"
	"net"
)

type IServer interface {
	Listen(network, address string) error
	Close() error
	Accept() (connection.Connection, error)
}

type Server struct {
	Listener net.Listener
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen(network, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	s.Listener = listener
	return nil
}

func (s *Server) Close() error {
	return s.Listener.Close()
}

func (s *Server) Accept() (*connection.Connection, error) {
	conn, err := s.Listener.Accept()
	if err != nil {
		return &connection.Connection{}, err
	}
	return connection.CreateConnection(conn), nil
}
