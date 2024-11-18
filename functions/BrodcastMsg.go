package functions

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// BrodcastMsg sends a message to all connected users except the sender.
// It also registers the message to the chat history for future reference.
func (c *Connections) BrodcastMsg(msg string, conn net.Conn) error {
	defer c.RegisterMsg(strings.Trim(msg, "\n"))
	defer c.Mutex.Unlock()

	c.Mutex.Lock()

	for userName, val := range c.Users {
		if val != conn {
			message := fmt.Sprintf("[%v][%v]:", time.Now().Format(time.DateTime), userName)
			val.Write([]byte(msg + message))
		}
	}

	return nil
}
