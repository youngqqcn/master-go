package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4}

	for x, n := range a {
		fmt.Printf("%d: %d\n", x, n)
		fmt.Printf("%p\n", &n)
	}
}
