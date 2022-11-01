package cborg

import (
	"errors"
	"net"

	"github.com/cborgdb/cborg-go-driver/pkg/cbor"
)

type Client struct {
	hostname string
	port     string
	conn     net.Conn
}

func NewClient(hostname string, port string) *Client {
	return &Client{hostname: hostname, port: port}
}

func (c *Client) Connect() error {
	var err error
	c.conn, err = net.Dial("tcp", c.hostname+":"+c.port)
	return err
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}

func (c *Client) Database(databaseName string) *Database {
	return &Database{databaseName, c}
}

func (c *Client) Databases() (*ResultListDatabases, error) {
	msg, _ := encodeMessage(OpListDatabases, nil, nil, nil, nil)
	c.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := c.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	// Get Message Length
	msgLength := cbor.DecodeUint(reply)
	s := cbor.DecodeTextString(reply[18:])
	headSizeS, _ := cbor.HeadSize(reply[18])
	var databases []string
	databases = append(databases, s)
	if s != "An error has occurred.\n" {
		// TODO build database array from cborseq textstrings
		for i := uint64(len(s) + headSizeS + 18); i < msgLength-1; i = i + uint64(len(s)+headSizeS) {
			s = cbor.DecodeTextString(reply[i:])
			headSizeS, _ = cbor.HeadSize(reply[18])
			databases = append(databases, s)
		}
		return &ResultListDatabases{true, "Success", databases}, nil
	} else {
		switch s {
		case "An error has occurred.\n":
			return &ResultListDatabases{false, s, []string{}}, errors.New("server an error has occurred")
		default:
			return &ResultListDatabases{false, s, []string{}}, errors.New("driver unknown reply")
		}
	}
}

func (c *Client) CreateDatabase(databaseName string) (*ResultCreateDatabase, error) {
	msg, _ := encodeMessage(OpCreateDatabase, &databaseName, nil, nil, nil)
	c.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := c.conn.Read(reply)
	if err != nil {
		return nil, err
	}
	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Database created.\n":
		return &ResultCreateDatabase{true, res, Database{databaseName, c}}, nil
	case res == "Database already exists.\n":
		return &ResultCreateDatabase{true, res, Database{databaseName, c}}, nil
	case res == "Database cannot created.\n":
		return &ResultCreateDatabase{false, res, Database{databaseName, c}}, errors.New("database cannot be created")
	default:
		return &ResultCreateDatabase{false, res, Database{databaseName, c}}, errors.New("driver unknown reply")
	}
}

func (c *Client) DropDatabase(databaseName string) (*ResultDropDatabase, error) {
	msg, _ := encodeMessage(OpDropDatabase, &databaseName, nil, nil, nil)
	c.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := c.conn.Read(reply)
	if err != nil {
		return nil, err
	}
	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Database deleted.\n":
		return &ResultDropDatabase{true, res}, nil
	case res == "Database not exists.\n":
		return &ResultDropDatabase{true, res}, nil
	case res == "Database cannot deleted.\n":
		return &ResultDropDatabase{false, res}, errors.New("database cannot be deleted")
	default:
		return &ResultDropDatabase{false, res}, errors.New("driver unknown reply")
	}
}
