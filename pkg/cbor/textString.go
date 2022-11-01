package cbor

func EncodeTextString(v string) []byte {
	return append(encodeUint(uint64(len(v)), 0x60), v...)
}

func DecodeTextString(b []byte) string {
	size, _ := HeadSize(b[0])
	len, _ := headGetUint(b)
	return string(b[size : uint64(size)+len])
}
