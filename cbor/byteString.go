package cbor

func EncodeByteString(v []byte) []byte {
	return append(encodeUint(uint64(len(v)), 0x40), v...)
}
