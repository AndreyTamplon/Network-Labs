package connection

import "net"

type IConnection interface {
	Connect(network, address string) error
	Close() error
	Read([]byte) (int, error)
	Write(data []byte) error
	RemoteAddr() net.Addr
}

type Connection struct {
	conn net.Conn
}

func NewConnection() *Connection {
	return &Connection{}
}

func CreateConnection(conn net.Conn) *Connection {
	return &Connection{
		conn: conn,
	}
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

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
