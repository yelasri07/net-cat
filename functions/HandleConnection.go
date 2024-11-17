package functions

import (
	"fmt"
	"net"
)

func (c *Connections) HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Welcome Message
	welcomeMessage := "Welcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    .       | ' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     -'       --'\n"
	_, err := conn.Write([]byte(welcomeMessage + "\n"))
	if err != nil {
		fmt.Println("write error:", err)
		return
	}

	// Handle User Name
	var userName string
	errName := c.handleName(conn, &userName)
	if errName != nil {
		return
	}

	// Send Previous Messages to New User
	if len(c.messages) > 1 {
		for _, message := range c.messages[1:] {
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("write error:", err)
				return
			}
		}
	}
	// Announce User Joining
	joiningMsg := fmt.Sprintf("\033[2K\r%s has joined our chat...\n", userName)
	errB := c.BrodcastMsg(joiningMsg, conn)
	if errB != nil {
		return
	}

	// Handle Incoming Messages
	errmsg := c.HandleMessage(conn, userName)
	if errmsg != nil {
		// Announce User Leaving
		leftMsg := fmt.Sprintf("\033[2K\r%s has left our chat...\n", userName)
		errB := c.BrodcastMsg(leftMsg, conn)
		c.RemoveClient(userName)
		if errB != nil {
			return
		}
	}
}
