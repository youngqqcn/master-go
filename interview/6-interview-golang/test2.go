package main

import "fmt"

func main() {

	{
		defer func() {
			fmt.Println("1111")
		}()
	}

	fmt.Println("22222")

}



func main() {

	func(){
		defer func() {
			fmt.Println("1111")
		}()
	}()

	fmt.Println("22222")

}
