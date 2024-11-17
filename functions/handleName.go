package functions

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func (c *Connections) handleName(conn net.Conn, userName *string) error {
	readChat := make([]byte, 4096)
	inputName := "[ENTER YOUR NAME]:"

check:
	_, err := conn.Write([]byte(inputName))
	if err != nil {
		return err
	}

	n, err := conn.Read(readChat)
	if err != nil {
		if err == io.EOF {
			fmt.Println("user out")
			return err
		}
		return err
	}

	if (*userName) == "" {
		if !ValidInput(readChat[:n]) {
			goto check
		}
		(*userName) = strings.Trim(string(readChat[:n]), "\n")
		c.AddClient((*userName), conn)
	}

	return nil
}
