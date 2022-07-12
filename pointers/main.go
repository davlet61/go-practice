package main

import "fmt"

func main() {
	x := 10
	y := &x

	fmt.Println(x, y)
	fmt.Printf("Type of X => %T\n", x)
	fmt.Printf("Pointer Type (y) -> %T\n", y)

	// Change the value of x through y
	*y = 20
	fmt.Println("new val for X", x)

	// get the val from the memory address
	fmt.Println("Value at memory address", *y)
}
