package functions

import (
	"net"
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

	if !ValidInput(readChat[:n-1]) || n == 1 {
		conn.Write([]byte("Invalid Input!!\n"))
		goto check
	}

	*userName = string(readChat[:n-1])

	return nil
}
