package main

import "fmt"

func main() {
	a := 5
	b := &a

	// b returns the location of a stored in the memory
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read value from the address
	fmt.Println(*b)

	// change value with pointer
	*b = 10
	fmt.Println(a)

}
