package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	time.Sleep(time.Second * 3)
	runtime.GC()
	fmt.Println("Done")
}
