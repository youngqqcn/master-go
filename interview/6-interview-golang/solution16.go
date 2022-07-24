package main

import "fmt"

// 下面代码输出什么?
func calc(index string, a, b *int) *int {
	ret := *a + *b
	fmt.Printf("%v,%v,%v,%v\n", index, *a, *b, ret)
	return &ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", &a, calc("10", &a, &b))
	a = 0
	defer calc("2", &a, calc("20", &a, &b))
	b = 1
}


// 10,1,2,3
// 20,0,2,2
// 2,0,2,2
// 1,0,3,3














// 10,1,2,3
// 20,0,2,2
// 2,0,2,2
// 1,1,3,4
