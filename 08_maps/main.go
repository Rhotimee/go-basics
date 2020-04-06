package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)

	// declare map and add key value
	data := map[string]int{"isiai": 23, "timi": 25, "isaiah": 30}

	// Assign
	emails["Bob"] = "bob@gmail.com"
	emails["timi"] = "timi@gmail.com"
	fmt.Println(emails)
	fmt.Println(emails["timi"])
	fmt.Println(len(emails))

	// delete
	delete(emails, "Bob")
	fmt.Println(emails)

	fmt.Println(data)

}
