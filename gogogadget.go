package gogogadget

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"runtime"
)

// BigEndian is used because I've been taught BigEndian all along and cannot really quickly parse LittleEndian byte ordering
var ByteOrdering binary.ByteOrder = binary.BigEndian
var NativeBitCount int

func init() {
	switch runtime.GOARCH {
	case "amd64":
		NativeBitCount = 64
	case "386":
		NativeBitCount = 32
	case "arm":
		NativeBitCount = 32 // I think?
	}
}

// BinaryRepresentation returns the binary representation as a string. 1234 in 32 bits would return 00000000 00000000 00000100 11010010
func BinaryRepresentation(x interface{}) string {
	buf := new(bytes.Buffer)

	var y interface{}
	switch x.(type) {
	case int:
		tmp, _ := x.(int)
		switch NativeBitCount {
		case 64:
			y = int64(tmp)
		case 32:
			y = int32(tmp)
		}
	case uint:
		tmp, _ := x.(uint)
		switch NativeBitCount {
		case 64:
			y = uint64(tmp)
		case 32:
			y = uint32(tmp)
		}
	default:
		y = x
	}

	err := binary.Write(buf, ByteOrdering, y)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	stringBuffer := ""
	for _, v := range buf.Bytes() {
		binaryRepresentation := fmt.Sprintf("%b",  v)
		difference := 8 - len(binaryRepresentation)
		prepend := ""
		if difference > 0 {
			prepend = strings.Repeat("0", difference)
		}
		stringBuffer += (prepend + binaryRepresentation + " ")
	}
	return stringBuffer
}