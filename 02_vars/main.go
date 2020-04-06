package main

import "fmt"

// List of the types
// https://tour.golang.org/basics/11

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
// 		// represents a Unicode code point

// float32 float64

// complex64 complex128
func main() {
	// var
	var name string = "Isaiah"
	var age = 26
	const isCool = true

	// shorthand
	email, pNumber := "yemi@n.co", "08083908"

	fmt.Println(name, age, isCool, email, pNumber)
	fmt.Printf("%T %T", name, age)

}
