package functions

import (
	"fmt"
	"time"
)

func (c *Connection) BroadCast(name, msg string) {
	c.Mutex.Lock()
	for nameClient, conn := range c.Clients {
		if nameClient != name {
			printMsg := fmt.Sprintf("\n[%s][%s]:", time.Now().Format(time.DateTime), nameClient)
			conn.Write([]byte(msg + printMsg))
		}
	}
	c.Mutex.Unlock()
}
