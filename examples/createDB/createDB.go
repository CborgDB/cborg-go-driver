package main

import (
	"os"

	"github.com/cborgdb/cborg-go-driver/pkg/ops"
)

func main() {
	// Connection
	conn, err := ops.Connect("127.0.0.1", "30000")
	if err != nil {
		os.Exit(1)
	}
	defer ops.Disconnect(conn)

	// Send request to create DB using Arg[1]
	ops.CreateDB(conn, os.Args[1])

	// Receive reply from the server
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	// Print reply from the server
	println("reply from server: ", string(reply[19:]))
}