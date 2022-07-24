package main

import (
	"fmt"
)

func main() {
	var x, y, sum_xy int

	fmt.Print("First Number: ")
	fmt.Scanln(&x)
	fmt.Print("Second Number: ")
	fmt.Scanln(&y)
	sum_xy = x + y
	fmt.Println("Sum: ", sum_xy)
}
