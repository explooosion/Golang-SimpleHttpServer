package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Print(add(1, 3))
}
