package cborg

import "github.com/cborgdb/cborg-go-driver/cbor"

func encodeMessage(op OpCode, dbName *string, collName *string, old []byte, new []byte) ([]byte, error) {
	msgWithoutLength := cbor.EncodeUint64(uint64(op))

	if dbName != nil {
		msgWithoutLength = append(msgWithoutLength, cbor.EncodeTextString(*dbName)...)
	}

	if collName != nil {
		msgWithoutLength = append(msgWithoutLength, cbor.EncodeTextString(*collName)...)
	}

	if old != nil {
		msgWithoutLength = append(msgWithoutLength, old...)
	}

	if new != nil {
		msgWithoutLength = append(msgWithoutLength, new...)
	}

	return append(cbor.EncodeUint64(uint64(len(msgWithoutLength))+uint64(9)), msgWithoutLength...), nil
}

func handleReply2() {

}
