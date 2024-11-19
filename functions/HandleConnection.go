package functions

import (
	"fmt"
	"net"
)

// HandleConnection manages the lifecycle of a client connection to the TCP chat server.
// This includes welcoming the client, handling their username, broadcasting messages,
// and managing their participation in the chat.
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
	conn.Write([]byte(welcomeMessage + "\n"))

	// Handle User Name
	var userName string
	errName := c.handleName(conn, &userName)
	if errName != nil {
		return
	}

	c.IncrementUserCount("+")

	if c.NbConn > 10 {
		conn.Write([]byte("Try logging in later, the chat is full.\n"))
		return
	}
	c.AddClient(userName, conn)

	// Send Previous Messages to New User
	if len(c.messages) > 1 {
		for _, message := range c.messages[1:] {
			conn.Write([]byte(message + "\n"))
		}
	}

	// Announce User Joining
	joiningMsg := fmt.Sprintf("\n%s has joined our chat...\n", userName)
	c.BrodcastMsg(joiningMsg, conn)

	// Handle Incoming Messages
	errmsg := c.HandleMessage(conn, userName)
	if errmsg != nil {
		// Announce User Leaving
		c.IncrementUserCount("-")
		leftMsg := fmt.Sprintf("\n%s has left our chat...\n", userName)
		c.BrodcastMsg(leftMsg, conn)


		c.RemoveClient(userName)
	}
}
