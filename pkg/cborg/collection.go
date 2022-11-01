package cborg

import (
	"errors"
	"fmt"

	"github.com/cborgdb/cborg-go-driver/pkg/cbor"
)

type Collection struct {
	name     string
	client   *Client
	database *Database
}

func (coll *Collection) Name() string {
	return coll.name
}

func (coll *Collection) Client() *Client {
	return coll.client
}

func (coll *Collection) Database() *Database {
	return coll.database
}

func (coll *Collection) InsertOne(item interface{}) (*ResultInsertOne, error) {
	msg, _ := encodeMessage(OpInsertOne, &coll.database.name, &coll.name, cbor.EncodeUint(uint64(item.(int))), nil)
	coll.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := coll.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Item inserted.\n":
		return &ResultInsertOne{true, res}, nil
	case res == "An error has occurred.\n":
		return &ResultInsertOne{false, res}, errors.New("an error has occurred")
	default:
		return &ResultInsertOne{false, res}, errors.New("driver unknown reply")
	}
}

func (coll *Collection) FindOne(query interface{}) (*ResultFindOne, error) {
	msg, _ := encodeMessage(OpFindOne, &coll.database.name, &coll.name, cbor.EncodeUint(uint64(query.(int))), nil)
	coll.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := coll.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Item found.\n":
		return &ResultFindOne{true, res, query}, nil
	case res == "Item not found.\n":
		return &ResultFindOne{false, res, query}, nil
	case res == "An error has occurred.\n":
		return &ResultFindOne{false, res, query}, errors.New("an error has occurred")
	default:
		return &ResultFindOne{false, res, query}, errors.New("driver unknown reply")
	}
}

func (coll *Collection) UpdateOne(old interface{}, new interface{}) (*ResultUpdateOne, error) {
	msg, _ := encodeMessage(OpUpdateOne, &coll.database.name, &coll.name, cbor.EncodeUint(uint64(old.(int))), cbor.EncodeUint(uint64(new.(int))))
	coll.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := coll.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Item updated.\n":
		return &ResultUpdateOne{true, res}, nil
	case res == "Item not found.\n":
		return &ResultUpdateOne{false, res}, nil
	case res == "An error has occurred.\n":
		return &ResultUpdateOne{false, res}, errors.New("an error has occurred")
	default:
		return &ResultUpdateOne{false, res}, errors.New("driver unknown reply")
	}
}

func (coll *Collection) UpdateAll(old interface{}, new interface{}) (*ResultUpdateAll, error) {
	fmt.Printf("cbor.EncodeUint(uint64(new.(int))): %x\n", cbor.EncodeUint(uint64(new.(int))))
	msg, _ := encodeMessage(OpUpdateAll, &coll.database.name, &coll.name, cbor.EncodeUint(uint64(old.(int))), cbor.EncodeUint(uint64(new.(int))))
	coll.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := coll.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Items updated.\n":
		return &ResultUpdateAll{true, res}, nil
	case res == "Item not found.\n":
		return &ResultUpdateAll{false, res}, nil
	case res == "An error has occurred.\n":
		return &ResultUpdateAll{false, res}, errors.New("an error has occurred")
	default:
		return &ResultUpdateAll{false, res}, errors.New("driver unknown reply")
	}
}

func (coll *Collection) DeleteOne(item interface{}) (*ResultDeleteOne, error) {
	msg, _ := encodeMessage(OpDeleteOne, &coll.database.name, &coll.name, cbor.EncodeUint(uint64(item.(int))), nil)
	coll.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := coll.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Item deleted.\n":
		return &ResultDeleteOne{true, res}, nil
	case res == "Item not found.\n":
		return &ResultDeleteOne{false, res}, nil
	case res == "An error has occurred.\n":
		return &ResultDeleteOne{false, res}, errors.New("an error has occurred")
	default:
		return &ResultDeleteOne{false, res}, errors.New("driver unknown reply")
	}
}

func (coll *Collection) DeleteAll(item interface{}) (*ResultDeleteAll, error) {
	msg, _ := encodeMessage(OpDeleteAll, &coll.database.name, &coll.name, cbor.EncodeUint(uint64(item.(int))), nil)
	coll.client.conn.Write(msg)

	// Receive reply
	reply := make([]byte, 1024)
	_, err := coll.client.conn.Read(reply)
	if err != nil {
		return nil, err
	}

	switch res := cbor.DecodeTextString(reply[18:]); {
	case res == "Items deleted.\n":
		return &ResultDeleteAll{true, res}, nil
	case res == "Item not found.\n":
		return &ResultDeleteAll{false, res}, nil
	case res == "An error has occurred.\n":
		return &ResultDeleteAll{false, res}, errors.New("an error has occurred")
	default:
		return &ResultDeleteAll{false, res}, errors.New("driver unknown reply")
	}
}
