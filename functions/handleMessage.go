package functions

import (
	"fmt"
	"net"
	"time"
)

// HandleMessage manages incoming messages from a specific client.
func (c *Connections) HandleMessage(conn net.Conn, userName string) error {
	readChat := make([]byte, 4096)

WriteAgain:
	now := time.Now()
	promptMessage := fmt.Sprintf("[%v][%v]:", now.Format(time.DateTime), userName)
	conn.Write([]byte(promptMessage))

	n, err := conn.Read(readChat)
	if err != nil {
		return err
	}

	if n == 1 || !ValidInput(readChat[:n-1]) {
		goto WriteAgain
	}

	broadcastMessage := fmt.Sprintf("\n[%v][%v]:%v", now.Format(time.DateTime), userName, string(readChat[:n]))

	c.BrodcastMsg(broadcastMessage, conn)

	goto WriteAgain
}
