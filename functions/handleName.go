package functions

import (
	"net"
	"strings"
)

// handleName prompts the user to enter a valid name and assigns it to userName.
func (c *Connections) handleName(conn net.Conn, userName *string, inputName string) error {
	readChat := make([]byte, 4096)

	for {
		conn.Write([]byte(inputName))

		n, err := conn.Read(readChat)
		if err != nil {
			return err
		}

		*userName = strings.TrimSpace(string(readChat[:n-1]))

		if !CheckSpaceName(*userName) || !ValidInput([]byte(*userName)) {
			conn.Write([]byte("Enter a name like : Ismail_Sayen | Youssef07 | !Mossab\n"))
			continue
		}

		if len(*userName) > 15 {
			conn.Write([]byte("The name should not exceed 15 letters.\n"))
			continue
		}

		if _, exist := c.Users[*userName]; exist {
			conn.Write([]byte("name already exist please try an other name.\n"))
			continue
		}
		return nil
	}
}
