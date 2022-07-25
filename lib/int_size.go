package lib

import (
	"fmt"
	"unsafe"
)

func GetSize() {
	i := 10
	size := unsafe.Sizeof(i)
	fmt.Println("The size of an int is: ", size)
}
