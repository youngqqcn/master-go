package main

import (
	"fmt"
	"sync"
)

/*
## 交替打印数字和字母

**问题描述**

使用两个 `goroutine` 交替打印序列，一个 `goroutine` 打印数字， 另外一个 `goroutine` 打印字母， 最终效果如下：

```bash
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
```
*/

func solution1() {

	chNum := make(chan int, 0)
	chChar := make(chan int, 0)

	go func() {
		for n := 1; n <= 28; {
			<-chNum
			fmt.Printf("%d", n)
			n++
			fmt.Printf("%d", n)
			n++
			chChar <- 0
		}
	}()

	chNum <- 0

	for c := 'A'; c <= 'Z'; {
		<-chChar
		fmt.Printf("%c", c)
		c++
		fmt.Printf("%c", c)
		c++
		chNum <- 0
	}
	<-chChar // 最后等数字打印完
	fmt.Println()
}

func solution() {
	chNum := make(chan int, 0)
	chChar := make(chan int, 0)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(w *sync.WaitGroup) {
		
		for n := 1; n <= 28; {
			select {
			case <-chNum:
				fmt.Printf("%d", n)
				n++
				fmt.Printf("%d", n)
				n++

				// 因为当n>=28时，打印字符的goroutine已经提前退出
				if n < 28 {
					chChar <- 0
				}
			}
		}
		wg.Done()
	}(wg)

	go func() {
		
		for c := 'A'; c <= 'Z'; {
			select {
			case <-chChar:
				fmt.Printf("%c", c)
				c++
				fmt.Printf("%c", c)
				c++
				chNum <- 0
			}
		}
	}()

	// 触发开始
	chNum <- 0

	wg.Wait()
	fmt.Println()
}

func main() {
	solution()
}
