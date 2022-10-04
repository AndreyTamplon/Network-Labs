package connection

import "net"

type IConnection interface {
	Connect(network, address string) error
	Close() error
	Read([]byte) (int, error)
	Write(data []byte) error
}

type Connection struct {
	conn net.Conn
}

func NewConnection() Connection {
	return Connection{}
}

func (c *Connection) Connect(network, address string) error {
	conn, err := net.Dial(network, address)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}

func (c *Connection) Read(data []byte) (int, error) {
	return c.conn.Read(data)
}

func (c *Connection) Write(data []byte) error {
	_, err := c.conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}
