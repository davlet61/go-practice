package main

import "fmt"

func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func squarer() func(int) int {
	return func(y int) int {
		return y * y
	}
}

func main() {
	mul3 := multiplier(3)
	mul5 := multiplier(5)
	squared := squarer()

	fmt.Println(mul3(2))
	fmt.Println(mul5(4))
	fmt.Println(squared(5))
}
