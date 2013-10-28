package gogogadget

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"runtime"
)

func init() {
	// BigEndian is used because I've been taught BigEndian all along and cannot really quickly parse LittleEndian byte ordering
	var ByteOrdering binary.ByteOrder = binary.BigEndian
	var NativeBitCount int
	if runtime.GOARCH == "amd64" {
		NativeBitCount = 64
	}

}

func Representation(x interface{}) string {
	buf := new(bytes.Buffer)

	var y interface{}
	switch x.(type) {
	case int:
		tmp, _ := x.(int)
		switch nativeBitCount {
		case 64:
			y = int64(tmp)
		case 32:
			y = int32(tmp)
		}
	case uint:
		tmp, _ := x.(uint)
		switch nativeBitCount {
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
	fmt.Println(newBuffer)
}