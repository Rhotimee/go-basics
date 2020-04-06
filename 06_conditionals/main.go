package main

import "fmt"

func main() {
	x := 10
	y := 10

	// if - else if - else
	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d \n", x, y)
	} else {
		fmt.Printf("%d is greater than %d \n", x, y)
	}

	color := "sd"

	// swich - case
	switch color {
	case "red":
		fmt.Printf("Color is red")
	case "blue":
		fmt.Printf("Color is blue")
	case "green":
		fmt.Printf("Color is green")
	default:
		fmt.Printf("I don't know the color")
	}
}
