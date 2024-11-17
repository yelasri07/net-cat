package functions

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func (c *Connections) BrodcastMsg(msg string, conn net.Conn) error {
	defer c.Mutex.Unlock()
	c.Mutex.Lock()
	for userName, val := range c.Users {
		if val != conn {
			message := fmt.Sprintf("[%v][%v]:", time.Now().Format(time.DateTime), userName)
			if _, err := val.Write([]byte(msg + message)); err != nil {
				return err
			}
		}
	}
	
	c.messages = append(c.messages, strings.Trim(msg, "\n"))
	fmt.Println("=>", c.messages)
	return nil
}
