package ops

import (
	"net"

	"github.com/cborgdb/cborg-go-driver/pkg/cbor"
)

func InsertOne(conn net.Conn, dbName string, collectionName string, item []byte) {
	operationEncoded := cbor.EncodeUint64(opInsertOne)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9 + len(item)))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	msg = append(msg, item...)
	conn.Write(msg)
}

func FindOne(conn net.Conn, dbName string, collectionName string, item []byte) {
	operationEncoded := cbor.EncodeUint64(opFindOne)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9 + len(item)))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	msg = append(msg, item...)
	conn.Write(msg)
}

func UpdateOne(conn net.Conn, dbName string, collectionName string, old []byte, new []byte) {
	operationEncoded := cbor.EncodeUint64(opUpdateOne)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9 + len(old) + len(new)))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	msg = append(msg, old...)
	msg = append(msg, new...)
	conn.Write(msg)
}

func UpdateAll(conn net.Conn, dbName string, collectionName string, old []byte, new []byte) {
	operationEncoded := cbor.EncodeUint64(opUpdateAll)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9 + len(old) + len(new)))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	msg = append(msg, old...)
	msg = append(msg, new...)
	conn.Write(msg)
}

func DeleteOne(conn net.Conn, dbName string, collectionName string, item []byte) {
	operationEncoded := cbor.EncodeUint64(opDeleteOne)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9 + len(item)))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	msg = append(msg, item...)
	conn.Write(msg)
}

func DeleteAll(conn net.Conn, dbName string, collectionName string, item []byte) {
	operationEncoded := cbor.EncodeUint64(opDeleteAll)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	collectionNameEncoded := cbor.EncodeTextString(collectionName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + len(collectionNameEncoded) + 9 + 9 + len(item)))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	msg = append(msg, collectionNameEncoded...)
	msg = append(msg, item...)
	conn.Write(msg)
}
