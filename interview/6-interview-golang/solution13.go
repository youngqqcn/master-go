package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 下面的代码会输出什么，并说明原因

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("i: %d\n", i )
			wg.Done()
		}()
	}

	wg.Wait()
}

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	wg := sync.WaitGroup{}
// 	wg.Add(3)
// 	for i := 0; i < 3; i++ {
// 		go func(i int) {
// 			fmt.Println("i: ", i)
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()
// }
