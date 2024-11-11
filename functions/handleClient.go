package functions

import (
	"fmt"
	"net"
)

func HandleClient(conn net.Conn) {
	r := make([]byte, 4096)
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

	conn.Write([]byte(welcomeMessage))

	conn.Write([]byte("[ENTER YOUR NAME]: "))

	n, _ := conn.Read(r)

	name := string(r[:n-1])

	fmt.Println(string(r))

	conn.Write([]byte("[2020-01-20 16:03:43][" + name + "]"))
}
