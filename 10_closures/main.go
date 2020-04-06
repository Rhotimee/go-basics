package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	i := 0
	sum := adder()
	for i < 10 {
		fmt.Println(sum(i))
		i++
	}
}


