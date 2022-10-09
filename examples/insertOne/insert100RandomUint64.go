package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/cborgdb/cborg-go-driver/pkg/cbor"
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
	reply := make([]byte, 4096)
	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	// Print reply from the server
	println("reply from server: ", string(reply[19:]))

	// Send request to create Collection Arg[2] in DB  Arg[1]
	ops.CreateCollection(conn, os.Args[1], os.Args[2])

	// Receive reply from the server
	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	// Print reply from the server
	println("reply from server: ", string(reply[18:]))

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		// Random Uint64
		var v uint64 = uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
		// Encode value
		encodedV := cbor.EncodeUint(v)
		// Insert
		ops.InsertOne(conn, os.Args[1], os.Args[2], encodedV)

		// Receive reply from the server
		reply := make([]byte, 4096)
		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		// Print reply from the server
		println("reply from server: ", string(reply[19:]))
	}

}
