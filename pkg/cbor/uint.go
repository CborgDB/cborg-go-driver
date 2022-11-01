package cbor

func EncodeUint8(v uint8) []byte {
	return encodeUint8(v, 0x00)
}

func EncodeUint16(v uint16) []byte {
	return encodeUint16(v, 0x00)
}

func EncodeUint32(v uint32) []byte {
	return encodeUint32(v, 0x00)
}

func EncodeUint64(v uint64) []byte {
	return encodeUint64(v, 0x00)
}

func EncodeUint(v uint64) []byte {
	return encodeUint(v, 0x00)
}

// TODO
func DecodeUint(e []byte) uint64 {
	v, _ := headGetUint(e)
	return v
}
