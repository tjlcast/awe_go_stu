package main

import "fmt"

func main() {
	var a, b = 1, 2
	a, b = b, a
	fmt.Printf("a:%d and b: %d\n", a, b)
}
