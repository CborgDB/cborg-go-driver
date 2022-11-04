package cborg

import (
	"errors"

	"github.com/cborgdb/cborg-go-driver/cbor"
)

type Database struct {
	name   string
	client *Client
}

func (db *Database) Client() *Client {
	return db.client
}

func (db *Database) Name() string {
	return db.name
}

func (db *Database) Collection(name string) *Collection {
	return &Collection{name, db.client, db}
}

func (db *Database) CreateCollection(name string) (*ResultCreateCollection, error) {
	msg, _ := encodeMessage(OpCreateCollection, &db.name, &name, nil, nil)
	db.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := db.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Collection created.\n":
		return &ResultCreateCollection{true, res, Collection{name, db.client, db}}, nil
	case res == "Collection already exists.\n":
		return &ResultCreateCollection{true, res, Collection{name, db.client, db}}, nil
	case res == "Collection cannot created.\n":
		return &ResultCreateCollection{false, res, Collection{name, db.client, db}}, errors.New("database cannot be created")
	default:
		return &ResultCreateCollection{false, res, Collection{name, db.client, db}}, errors.New("driver unknown reply")
	}
}

func (db *Database) DropCollection(name string) (*ResultDropCollection, error) {
	msg, _ := encodeMessage(OpDropCollection, &db.name, &name, nil, nil)
	db.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := db.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Collection deleted.\n":
		return &ResultDropCollection{true, res}, nil
	case res == "Collection not exists.\n":
		return &ResultDropCollection{true, res}, nil
	case res == "Collection cannot deleted.\n":
		return &ResultDropCollection{false, res}, errors.New("collection cannot be deleted")
	default:
		return &ResultDropCollection{false, res}, errors.New("driver unknown reply")
	}
}

func (db *Database) Collections() (*ResultListCollections, error) {
	msg, _ := encodeMessage(OpListCollections, &db.name, nil, nil, nil)
	db.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := db.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	// Get Message Length
	msgLength := cbor.DecodeUint(reply)
	s := cbor.DecodeTextString(reply[18:])
	headSizeS, _ := cbor.HeadSize(reply[18])
	var collections []string
	collections = append(collections, s)
	if s != "An error has occurred.\n" {
		for i := uint64(len(s) + headSizeS + 18); i < msgLength-1; i = i + uint64(len(s)+headSizeS) {
			s = cbor.DecodeTextString(reply[i:])
			headSizeS, _ = cbor.HeadSize(reply[18])
			collections = append(collections, s)
		}
		return &ResultListCollections{true, "Success", collections}, nil
	} else {
		switch s {
		case "An error has occurred.\n":
			return &ResultListCollections{false, s, []string{}}, errors.New("server an error has occurred")
		default:
			return &ResultListCollections{false, s, []string{}}, errors.New("driver unknown reply")
		}
	}
}
