package functions

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
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

// Logs adds a message to the chat history.
func (c *Connections) Logs(s string) {
	defer c.Unlock()
	c.Lock()
	file, err := os.OpenFile("./logs/activities.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	defer file.Close()
	msg := fmt.Sprintf("%v : %v\n", time.Now().Format(time.DateTime), strings.Trim(s, "\n"))
	_, err = file.WriteString(msg)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
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
