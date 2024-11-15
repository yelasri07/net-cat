package functions

import (
	"fmt"
	"time"
)

func (c *Connection) BroadCast(name, msg string) {
	c.Mutex.Lock()
	for nameClient, conn := range c.Clients {
		if nameClient != name {
			printMsg := fmt.Sprintf("\n[%v][%v]: %v\n", time.Now().Format(time.DateTime), name, msg)
			conn.Write([]byte(printMsg))
		}
	}
	c.Mutex.Unlock()
}
