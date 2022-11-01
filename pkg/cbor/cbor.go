package cbor

import (
	"encoding/binary"
	"errors"
	"math"
)

func encodeUint8(v uint8, offset byte) []byte {
	if v < 24 {
		var encodedV = make([]byte, 1)
		encodedV[0] = offset | byte(v)
		return encodedV
	} else {
		var encodedV = make([]byte, 2)
		encodedV[0] = offset | byte(0x18)
		encodedV[1] = byte(v)
		return encodedV
	}
}

func encodeUint16(v uint16, offset byte) []byte {
	var encodedV = make([]byte, 3)
	encodedV[0] = offset | byte(0x19)
	binary.BigEndian.PutUint16(encodedV[1:], v)
	return encodedV
}

func encodeUint32(v uint32, offset byte) []byte {
	var encodedV = make([]byte, 5)
	encodedV[0] = offset | byte(0x1A)
	binary.BigEndian.PutUint32(encodedV[1:], v)
	return encodedV
}

func encodeUint64(v uint64, offset byte) []byte {
	var encodedV = make([]byte, 9)
	encodedV[0] = offset | byte(0x1B)
	binary.BigEndian.PutUint64(encodedV[1:], v)
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

func HeadSize(ib byte) (int, error) {
	switch size := ib & 0b000_11111; {
	case size < 24:
		return 1, nil
	case size == 24:
		return 2, nil
	case size == 25:
		return 3, nil
	case size == 26:
		return 5, nil
	case size == 27:
		return 9, nil
	default:
		return -1, errors.New("cbor cannot get size of uint")
	}
}

func headGetUint(b []byte) (uint64, error) {
	size := b[0] & 0b000_11111
	switch {
	case size < byte(24):
		return uint64(b[0] & 0b000_11111), nil
	case size == byte(24):
		return uint64(b[1]), nil
	case size == byte(25):
		return uint64(binary.BigEndian.Uint16(b[1:])), nil
	case size == byte(26):
		return uint64(binary.BigEndian.Uint32(b[1:])), nil
	case size == byte(27):
		return binary.BigEndian.Uint64(b[1:]), nil
	default:
		return 0, errors.New("cbor cannot get uint")
	}
}
