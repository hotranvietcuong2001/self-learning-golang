// The package “main” tells the Go compiler that the package should compile as
// an executable program instead of a shared library
package main

// import the package “fmt” for using "Println"
import (
	"fmt"
)

func main() { // The left curly bracket { cannot come at the start of a line
	fmt.Println("Hello World")
}
