package ops

import (
	"net"

	"github.com/cborgdb/cborg-go-driver/pkg/cbor"
)

func CreateDB(conn net.Conn, dbName string) {
	operationEncoded := cbor.EncodeUint64(opCreateDB)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + 9 + 9))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	conn.Write(msg)
}

func DropDB(conn net.Conn, dbName string) {
	operationEncoded := cbor.EncodeUint64(opDropDB)
	dbNameEncoded := cbor.EncodeTextString(dbName)
	msgLength := cbor.EncodeUint64(uint64(len(dbNameEncoded) + 9 + 9))

	msg := append(msgLength, operationEncoded...)
	msg = append(msg, dbNameEncoded...)
	conn.Write(msg)
}

func ListDBs(conn net.Conn) {
	operationEncoded := cbor.EncodeUint64(opListDBs)
	msgLength := cbor.EncodeUint64(uint64(9 + 9))

	msg := append(msgLength, operationEncoded...)
	conn.Write(msg)
}
