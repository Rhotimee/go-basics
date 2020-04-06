package main

import "fmt"

func main() {
	// For is Go's "while"
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// normal for
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// range
	ids := []int{33, 12, 12, 41, 1, 234, 5}

	// var sum int // default value is 0. So, this will work.
	sum := 0
	for i, id := range ids {
		println(i, id)
		sum += id
	}

	println("sum", sum)

	// range with map
	data := map[string]int{"isiai": 23, "timi": 25, "isaiah": 30}

	for k, v := range data {
		fmt.Printf("%s has a value of %d \n", k, v)
	}

}
