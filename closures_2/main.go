package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	next := counter()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
