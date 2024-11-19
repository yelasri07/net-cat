package functions

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// HandleMessage manages incoming messages from a specific client.
func (c *Connections) HandleMessage(conn net.Conn, userName string) error {
	readChat := make([]byte, 4096)
	
	for {
		now := time.Now()
		promptMessage := fmt.Sprintf("[%v][%v]:", now.Format(time.DateTime), userName)
		conn.Write([]byte(promptMessage))

		n, err := conn.Read(readChat)
		if err != nil {
			return err
		}

		trimMsg := strings.TrimSpace(string(readChat[:n-1]))

		if !ValidInput([]byte(trimMsg)) {
			continue
		}

		// Handle user name change.
		if trimMsg == "/changename" {
			var newName string
			err := c.handleName(conn, &newName, "[ENTER YOUR NEW NAME]: ")
			if err != nil {
				continue
			}
			userName = c.ChangeName(conn, newName, userName)
			continue
		}

		broadcastMessage := fmt.Sprintf("\n[%v][%v]:%v", now.Format(time.DateTime), userName, string(trimMsg))
		c.BrodcastMsg(broadcastMessage, conn)
	}
}
