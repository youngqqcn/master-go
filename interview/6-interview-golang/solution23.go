package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 在 golang 协程和channel配合使用
// 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
// 另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

func solution() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int, 0)
	// chStop := make(chan bool, 0)

	// producer
	go func() {
		defer wg.Done()
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 5; i++ {
			ch <- rnd.Int() % 100000
		}
		// chStop <- true
		close(ch)
	}()

	// consumer
	go func() {
		// for {
		// 	select {
		// 	case n := <- ch:
		// 		fmt.Printf("%d\n", n)
		// 	case <-chStop:
		// 		wg.Done()
		// 		return
		// 	}
		// }
		defer wg.Done()

		// 可以用`for`循环来读取数据，当管道关闭后，`for` 退出
		for n := range ch {
			fmt.Printf("%d\n", n)
		}
		
	}()

	wg.Wait()
}

func main() {

	solution()

}
