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

WriteAgain:
	now := time.Now()
	promptMessage := fmt.Sprintf("[%v][%v]:", now.Format(time.DateTime), userName)
	conn.Write([]byte(promptMessage))

	n, err := conn.Read(readChat)
	if err != nil {
		return err
	}

	trimMsg := strings.TrimSpace(string(readChat[:n-1]))

	if !ValidInput([]byte(trimMsg)) {
		goto WriteAgain
	}

	// Handle user name change.
	if strings.HasPrefix(string(trimMsg), "--rename:") {
		newName := strings.TrimSpace(string(trimMsg[9:]))
		if !CheckSpaceName(newName) {
			conn.Write([]byte("Enter a name like : Ismail_Sayen | Youssef07 | !Mossab\n"))
			goto WriteAgain
		} else {
			userName = c.ChangeName(conn, newName, userName)
			goto WriteAgain
		}
	}

	broadcastMessage := fmt.Sprintf("\n[%v][%v]:%v", now.Format(time.DateTime), userName, string(trimMsg))

	c.BrodcastMsg(broadcastMessage, conn)

	goto WriteAgain
}
