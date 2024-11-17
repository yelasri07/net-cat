package functions

import (
	"net"
	"sync"
)

type Connections struct {
	Users    map[string]net.Conn
	messages []string
	sync.Mutex
}

func NewConnection() *Connections {
	return &Connections{
		Users: make(map[string]net.Conn),
	}
}

func (c *Connections) AddClient(name string, conn net.Conn) {
	c.Mutex.Lock()
	c.Users[name] = conn
	c.Mutex.Unlock()
}

func (c *Connections) RemoveClient(name string) {
	c.Mutex.Lock()
	delete(c.Users, name)
	c.Mutex.Unlock()
}

func (c *Connections) RegisterMsg(s string) {
	c.Mutex.Lock()
	c.messages = append(c.messages, s)
	c.Mutex.Unlock()
}
