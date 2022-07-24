package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	if stu == nil {
		fmt.Println("C")
	}
	return stu
}

func main() {
	p := live()
	if p == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}
