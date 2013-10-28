package main

import(
	"github.com/chewxy/gogogadget"
	"flag"
	"strconv"
)
var endian = flag.String("e", "", "Endianness: b or s")
var t = flag.String("t", "", "Type: all basic numeric Go types (int, byte, uint, float, etc)")
var i = flag.String("i", "", "Input")

func main() {
	flag.Parse()
	if *t == "" {
		panic("Type needs to be passed in")
	}
	if *i == "" {
		panic("You need to pass in an input")
	}

	var v interface{}

	switch *t {
	case "int":
		v = strconv.ParseInt(*t, 10, 0)
	case "int8":
		v = int8(strconv.ParseInt(*t, 10, 0))
	case "int16":
		v = int16(strconv.ParseInt(*t, 10, 0))
	case "int32":
		v = int32(strconv.ParseInt(*t, 10, 0))
	case "rune":
		v = rune(strconv.ParseInt(*t, 10, 0))
	case "int64":
		v = int64(strconv.ParseInt(*t, 10, 0))
	case "uint":
		v = strconv.ParseUint(*t, 10, 0)
	case "uint8":
		v = uint8(strconv.ParseUint(*t, 10, 0))
	case "byte":
		v = byte(strconv.ParseUint(*t, 10, 0))
	case "uint16":
		v = uint16(strconv.ParseUint(*t, 10, 0))
	case "uint32":
		v = uint32(strconv.ParseUint(*t, 10, 0))
	case "uint64":
		v = uint64(strconv.ParseUint(*t, 10, 0))

	}
}