package functions

import (
	"net"
	"sync"
)

type Connection struct {
	Clients map[string]net.Conn
	Mutex   sync.Mutex
}

func NewConnection() *Connection {
	return &Connection{
		Clients: make(map[string]net.Conn),
	}
}

func (c *Connection) AddClient(name string, conn net.Conn) {
	c.Mutex.Lock()
	c.Clients[name] = conn
	c.Mutex.Unlock()
}

func (c *Connection) RemoveClient(name string) {
	c.Mutex.Lock()
	delete(c.Clients, name)
	c.Mutex.Unlock()
}
