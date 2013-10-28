gogogadget
==========

GoGoGadget is a collection of useful tools for introspection and inspection in Go. Inspector Gadget says "Go Go Gadget" a lot. 

Currently the only inspection capability is to inspect the binary representation of various basic types in Go.

Example:
```Go
import (
	"github.com/chewxy/gogogadget"
	"fmt"
)

func main() {
	x := 1234
	fmt.Println(gogogadget.BinaryRepresentation(x))
}
```

On a 64-bit machine, this will yield (in BigEndian byte ordering):

```
00000000 00000000 00000000 00000000 00000000 00000000 00000100 11010010
```

On a 32-bit machine this will yield:

```
00000000 00000000 00000100 11010010
```

###On Endianness###

The reason why by default it is BigEndian is because that's the way binary was taught in high school. I never got used to LittleEndian notation and could not read it intuitively. Big Endian representation is in my opinion also a lot more intuitive when considering bitwise operations.

However, you can change the endianness yourself. Simply set `gogogadget.ByteOrdering = binary.LittleEndian` and you're good to go


#Why?#

Why was this package written? In bid to practice some bitwise math, I needed to visualize what I was doing, and occasionally even manually do it on pencil and paper. In this author's opinion, this is also useful for other forms of learning. More inspection methods will be added for even more usefulness in the future.