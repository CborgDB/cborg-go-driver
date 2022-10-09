package cbor

import "math"

func encodeUint8(v uint8, offset byte) []byte {
	if v < 24 {
		var encodedV = make([]byte, 1)
		encodedV[0] = byte(offset | v)
		return encodedV
	} else {
		var encodedV = make([]byte, 2)
		encodedV[0] = byte(offset | 0x18)
		encodedV[1] = byte(v)
		return encodedV
	}
}

func encodeUint16(v uint16, offset byte) []byte {
	var encodedV = make([]byte, 3)
	encodedV[0] = byte(0x19)
	encodedV[1] = byte((v >> 8) & 0xff)
	encodedV[2] = byte(v & 0xff)
	return encodedV
}

func encodeUint32(v uint32, offset byte) []byte {
	var encodedV = make([]byte, 5)
	encodedV[0] = byte(0x1A)
	encodedV[1] = byte((v >> 24) & 0xff)
	encodedV[2] = byte((v >> 16) & 0xff)
	encodedV[3] = byte((v >> 8) & 0xff)
	encodedV[4] = byte(v & 0xff)
	return encodedV
}

func encodeUint64(v uint64, offset byte) []byte {
	var encodedV = make([]byte, 9)
	encodedV[0] = byte(0x1B)
	encodedV[1] = byte((v >> 56) & 0xff)
	encodedV[2] = byte((v >> 48) & 0xff)
	encodedV[3] = byte((v >> 40) & 0xff)
	encodedV[4] = byte((v >> 32) & 0xff)
	encodedV[5] = byte((v >> 24) & 0xff)
	encodedV[6] = byte((v >> 16) & 0xff)
	encodedV[7] = byte((v >> 8) & 0xff)
	encodedV[8] = byte(v & 0xff)
	return encodedV
}

func encodeUint(v uint64, offset byte) []byte {
	if v <= math.MaxUint8 {
		return encodeUint8(uint8(v), offset)
	} else if v <= math.MaxUint16 {
		return encodeUint16(uint16(v), offset)
	} else if v <= math.MaxUint32 {
		return encodeUint32(uint32(v), offset)
	} else {
		return encodeUint64(v, offset)
	}
}
