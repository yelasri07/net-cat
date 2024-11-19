package functions

import (
	"fmt"
	"net"
)

// ChangeName updates the user's name.
func (c *Connections) ChangeName(conn net.Conn, newName, oldName string) string {
	c.RemoveClient(oldName)
	c.AddClient(newName, conn)
	alertChangeName := fmt.Sprintf("\n%s has changed their name to %s", oldName, newName)
	c.BrodcastMsg(alertChangeName, conn)
	c.Logs(alertChangeName)
	return newName
}
