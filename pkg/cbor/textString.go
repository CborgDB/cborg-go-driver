package cbor

func EncodeTextString(v string) []byte {
	lengthEncoded := encodeUint(uint64(len(v)), 0x60)
	encodedTextString := append(lengthEncoded, v...)
	return encodedTextString
}
