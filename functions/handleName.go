package functions

import (
	"net"
	"strings"
)

// handleName prompts the user to enter a valid name and assigns it to userName.
func (c *Connections) handleName(conn net.Conn, userName *string) error {
	readChat := make([]byte, 4096)
	inputName := "[ENTER YOUR NAME]: "

check:

	conn.Write([]byte(inputName))

	n, err := conn.Read(readChat)
	if err != nil {
		return err
	}

	trimUserName := strings.TrimSpace(string(readChat[:n-1]))

	if !CheckSpaceName(trimUserName) || !ValidInput([]byte(trimUserName)) {
		conn.Write([]byte("Enter a name like : Ismail_Sayen | Youssef07 | !Mossab\n"))
		goto check
	}

	(*userName) = string(trimUserName)
	if _, exist := c.Users[(*userName)]; exist {
		conn.Write([]byte("name already exist please try an other name.\n"))
		goto check
	}
	return nil
}
