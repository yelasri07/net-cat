package functions

import (
	"fmt"
	"net"
	"time"
)

func (c *Connections) HandleMessage(conn net.Conn, userName string) error {
	readChat := make([]byte, 4096)
WriteAgain:
	now := time.Now()
	message := fmt.Sprintf("[%v][%v]:", now.Format(time.DateTime), userName)
	_, errw := conn.Write([]byte(message))
	n, err2 := conn.Read(readChat)
	if err2 != nil {
		return err2
	}
	if errw != nil {
		return errw
	}

	if !ValidInput(readChat[:n]) {
		goto WriteAgain
	}

	message2 := fmt.Sprintf("\n[%v][%v]:%v", now.Format(time.DateTime), userName, string(readChat[:n]))

	errB := c.BrodcastMsg(message2, conn)

	if errB != nil {
		return errB
	}
	goto WriteAgain
}
