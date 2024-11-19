package functions

import (
	"net"
	"sync"
)

// Connections manages active users and chat messages.
type Connections struct {
	Users    map[string]net.Conn
	messages []string
	NbConn   int
	sync.Mutex
}

// NewConnection initializes a new Connections instance.
func NewConnection() *Connections {
	return &Connections{
		Users: make(map[string]net.Conn),
	}
}

// AddClient adds a new user to the active users map.
func (c *Connections) AddClient(name string, conn net.Conn) {
	c.Lock()
	c.Users[name] = conn
	c.Unlock()
}

// RemoveClient removes a user from the active users map.
func (c *Connections) RemoveClient(name string) {
	c.Lock()
	delete(c.Users, name)
	c.Unlock()
}

// RegisterMsg adds a message to the chat history.
func (c *Connections) RegisterMsg(s string) {
	c.Lock()
	c.messages = append(c.messages, s)
	c.Unlock()
}

func (c *Connections) IncrementUserCount(operation string) {
	c.Lock()
	switch operation {
	case "+":
		c.NbConn++
	case "-":
		c.NbConn--
	}
	c.Unlock()
}
