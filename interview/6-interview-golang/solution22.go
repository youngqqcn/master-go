package main

import "fmt"

func main() {
	var p interface{} = nil
	if p == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	var q interface{} = (*interface{})(nil)
	if q == nil {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

}
