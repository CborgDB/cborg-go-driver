package ops

import (
	"fmt"
	"net"
	"os"
)

func Connect(hostname string, port string) (net.Conn, error) {
	c, err := net.Dial("tcp", hostname+":"+port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR]: ", err.Error())
	}
	return c, err
}

func Disconnect(conn net.Conn) {
	conn.Close()
}
