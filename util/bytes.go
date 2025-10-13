package util

import (
	"fmt"
)

const (
	AsciiComma = 0x2c
	AsciiNewLine = 0xa
)

func BoolToByteSlice(b bool) []byte {
	bytes := make([]byte, 1)

	if b {
		bytes[0] = 1
	} else {
		bytes[0] = 0
	}

	return bytes
}

func BytesSliceToBool(bytes []byte) (bool, error) {
	if len(bytes) == 0 {
		return false, fmt.Errorf("Can't convert an emtpy bytes slice to a bool")
	}

	if bytes[0] == 0 {
		return false, nil
	}

	return true, nil
}
