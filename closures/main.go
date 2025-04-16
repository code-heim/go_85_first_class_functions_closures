package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println("Double of 5 =", double(5))
	fmt.Println("Triple of 5 =", triple(5))
}
