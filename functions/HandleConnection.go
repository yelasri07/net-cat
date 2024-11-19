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
	if c.NbConn > 9 {
		conn.Write([]byte("Try logging in later, the chat is full.\n"))
		return
	}

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

	c.AddClient(userName, conn)

	// Send Previous Messages to New User
	if len(c.messages) > 1 {
		for _, message := range c.messages[1:] {
			conn.Write([]byte(message + "\n"))
		}
	}

	// Announce User Joining
	joiningMsg := fmt.Sprintf("\n%s has joined our chat...", userName)
	c.BrodcastMsg(joiningMsg, conn)
	c.Logs(joiningMsg)

	// Handle Incoming Messages
	errmsg := c.HandleMessage(conn, userName)
	if errmsg != nil {
		
		// Announce User Leaving
		c.IncrementUserCount("-")
		leftMsg := fmt.Sprintf("\n%s has left our chat...", userName)
		c.BrodcastMsg(leftMsg, conn)
		c.Logs(leftMsg)
		c.RemoveClient(userName)
	}
}
