package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota
	d    = iota
)

const a0 = iota // a0 = 0  // const出现, iota初始化为0

const (
    a1 = iota   // a1 = 0   // 又一个const出现, iota初始化为0
    a2 = iota   // a1 = 1   // const新增一行, iota 加1
    a3 = 6      // a3 = 6   // 自定义一个常量
    a4          // a4 = 6   // 不赋值就和上一行相同
    a5 = iota   // a5 = 4   // const已经新增了4行, 所以这里是4
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
}

/*
使用的规则如下:

每当const出现时, 都会使iota初始化为0.
const中每新增一行常量声明将使iota计数一次.

*/

