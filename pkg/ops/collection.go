package ops

import (
	"net"

	"github.com/cborgdb/cborg-go-driver/pkg/cbor"
)

func CreateCollection(conn net.Conn, dbName string, collectionName string) {
	operationEncoded := cbor.EncodeUint64(opCreateCollection)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	conn.Write(msg)
}

func DropCollection(conn net.Conn, dbName string, collectionName string) {
	operationEncoded := cbor.EncodeUint64(opDropCollection)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	conn.Write(msg)
}

func ListCollections(conn net.Conn, dbName string) {
	operationEncoded := cbor.EncodeUint64(opListCollections)
	msgLength := cbor.EncodeUint64(uint64(9 + 9))
	dbNameEncoded := cbor.EncodeTextString(dbName)

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	conn.Write(msg)
}
