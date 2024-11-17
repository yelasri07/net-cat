package functions

import (
	"fmt"
	"net"
	"time"
)

func (c *Connection) HandleClient(conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	readChat := make([]byte, 4096)
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

	c.Mutex.Lock()
	conn.Write([]byte(welcomeMessage))
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	c.Mutex.Unlock()

	n, err := conn.Read(readChat)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	if !isValidInput(readChat, n) {
		conn.Write([]byte("Invalid name...\n"))
		return
	}

	name := string(readChat[:n-1])
	c.AddClient(name, conn)

	joinClient := fmt.Sprintf("\n%s has joined our chat...", name)
	c.BroadCast(name, joinClient)

	defer func() {
		leftClient := fmt.Sprintf("\n%s has left our chat...", name)
		c.BroadCast(name, leftClient)
		c.RemoveClient(name)
	}()

	for {
		printMsg := fmt.Sprintf("[%v][%v]:", time.Now().Format(time.DateTime), name)
		c.Mutex.Lock()
		conn.Write([]byte(printMsg))
		c.Mutex.Unlock()
		n, err := conn.Read(readChat)
		if err != nil || n < 1 {
			return
		}

		msg := string(readChat[:n-1])
		msg = fmt.Sprintf("\n[%v][%v]:%v", time.Now().Format(time.DateTime), name, msg)
		c.BroadCast(name, msg)
	}
}
