package main

import (
	"fmt"
)

//go:export sum
func sum(x int32, y int32) int32 {
	return x + y
}

//go:export sum2
func sum2(
	x int32, y int32) int32 {
	return x + y
}

func main() {
	fmt.Println("wasm-demo-api plugin.")
}
