package main

import (
	"fmt"
)

func main() {
	var last_name, first_name string
	fmt.Print("Input your last name: ")
	fmt.Scanln(&last_name)
	fmt.Print("Input you first name: ")
	fmt.Scanln(&first_name)
	fmt.Printf("Welcome %s %s\n", last_name, first_name)
}
