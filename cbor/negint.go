package cbor

func EncodeNegint8(v uint8) []byte {
	return encodeUint8(v, 0x20)
}

func EncodeNegint16(v uint16) []byte {
	return encodeUint16(v, 0x20)
}

func EncodeNegint32(v uint32) []byte {
	return encodeUint32(v, 0x20)
}

func EncodeNegint64(v uint64) []byte {
	return encodeUint64(v, 0x20)
}
