package main

import "fmt"

func main() {
	//  Arrays

	// declear array with capacity of 2
	var fruitArr [2]string

	// assign array with indexes
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// declare and assign array with values bounded. Array capacity of 5.
	names := [5]int{1, 2, 3, 4, 5}

	// dynamic array
	year := []int{1999, 2000, 2001, 2002, 2003, 2004}

	fmt.Println(fruitArr)
	fmt.Println(names)
	// cap and len is to get the capacity and length of the array
	fmt.Println(year, cap(year), len(year))
	// get range
	fmt.Println(year[1:3])

}
