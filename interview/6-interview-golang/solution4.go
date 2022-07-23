package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000) 
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}


// ch 是容量为1000的channel， 所以第一个goroutine写入的时候不会阻塞，当第一个goroutine写完之后马上退出，close(ch)关闭了ch
// 此时，第2个goroutine还在运行，往ch写入数据, 会触发panic
