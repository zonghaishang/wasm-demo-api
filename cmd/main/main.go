package main

import (
	"fmt"
	"unsafe"
)

//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

//export sum2
func sum2(x int32, y int32) int32 {
	return x + y
}

func main() {
	fmt.Println("wasm-demo-api plugin.")
}
