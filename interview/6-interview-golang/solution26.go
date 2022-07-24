// ## 1. 写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出?

package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go func() {
// 		// 1 在这里需要你写算法
// 		// 2 要求每秒钟调用一次proc函数

// 		// 3 要求程序不能退出

// 		defer func() {
// 			if r := recover(); r != nil {
// 				// fmt.Printf("%s\n", r)
// 			}
// 			main()
// 		}()

// 		// for {
// 		// t := time.NewTicker(time.Nanosecond)
// 		// select {
// 		// case <-t.C:
// 		proc()
// 		// }
// 		// }

// 	}()

// 	select {}
// }

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		for {
			t := time.NewTicker(time.Second)
			select {
			case <-t.C:
				go func() {
					defer func() {
						if r := recover(); r != nil {
							fmt.Printf("%s\n", r)
						}
					}()

					proc()
				}()
			}
		}

	}()

	select {}
}

func proc() {
	panic("ok")
}
