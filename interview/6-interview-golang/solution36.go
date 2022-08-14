package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	if ch == nil {
		fmt.Println("ch is nil")
	}
	// ch := make(chan int, 0)
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		// time.Sleep(time.Millisecond *10)
		// close(ch)
		// fmt.Println("closed")
	}()
	<-ch
	time.Sleep(time.Second)
	fmt.Println(count)
}
