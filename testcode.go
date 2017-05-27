package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("the var is %d ",add(42, 13))
}