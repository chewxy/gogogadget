package main

import(
	"github.com/chewxy/gogogadget"
	"flag"
	"strconv"
	"fmt"
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
	var tmp interface{}
	var err error

	switch *t {
	case "int":
		tmp, err = strconv.ParseInt(*i, 10, 0)
		v = tmp
	case "int8":
		tmp, err = strconv.ParseInt(*i, 10, 8)
		tmpInt, _ := tmp.(int64)
		v = int8(tmpInt)
	case "int16":
		tmp, err = strconv.ParseInt(*i, 10, 16)
		tmpInt, _ := tmp.(int64)
		v = int16(tmpInt)
	case "int32":
		tmp, err = strconv.ParseInt(*i, 10, 32)
		tmpInt, _ := tmp.(int64)
		v = int32(tmpInt)
	case "rune":
		tmp, err = strconv.ParseInt(*i, 10, 32)
		tmpInt, _ := tmp.(int64)
		v = rune(tmpInt)
	case "int64":
		tmp, err = strconv.ParseInt(*i, 10, 64)
		tmpInt, _ := tmp.(int64)
		v = int64(tmpInt)
	case "uint":
		tmp, err = strconv.ParseUint(*i, 10, 0)
		v = tmp
	case "uint8":
		tmp, err = strconv.ParseUint(*i, 10, 8)
		tmpUint, _ := tmp.(uint64)
		v = uint8(tmpUint)
	case "byte":
		tmp, err = strconv.ParseUint(*i, 10, 8)
		tmpUint, _ := tmp.(uint64)
		v = byte(tmpUint)
	case "uint16":
		tmp, err = strconv.ParseUint(*i, 10, 16)
		tmpUint, _ := tmp.(uint64)
		v = uint16(tmpUint)
	case "uint32":
		tmp, err = strconv.ParseUint(*i, 10, 32)
		tmpUint, _ := tmp.(uint64)
		v = uint32(tmpUint)
	case "uint64":
		tmp, err = strconv.ParseUint(*i, 10, 64)
		tmpUint, _ := tmp.(uint64)
		v = uint64(tmpUint)
	case "float32":
		tmp, err = strconv.ParseFloat(*i, 32)
		tmpFloat, _ := tmp.(float64)
		v = float32(tmpFloat)
	case "float64":
		tmp, err = strconv.ParseFloat(*i, 64)
		v = tmp
	}
	if err != nil {
		panic(fmt.Sprintf("Error: %s\n", err))
	}
	fmt.Println(*i + "(" + *t + "): " + gogogadget.BinaryRepresentation(v))
}