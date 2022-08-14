package main

import (
	"fmt"
	"net/http"
	"runtime/pprof"
	"sync"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func main() {

	go func() {
		wg := sync.WaitGroup{}
		c := make(chan struct{})
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(num int, close <-chan struct{}) {
				defer wg.Done() // 如果没有这一句，则需要在WaitTimeout中考虑wg.Wait()退出的问题
				<-close
				fmt.Println(num)
			}(i, c)
		}

		if WaitTimeout(&wg, time.Second*5) {
			close(c)
			fmt.Println("timeout exit")
		}
		time.Sleep(time.Second * 10)
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":11181", nil)
}

// func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
// 	// 要求手写代码
// 	// 要求sync.WaitGroup支持timeout功能
// 	// 如果timeout到了超时时间返回true
// 	// 如果WaitGroup自然结束返回false

// 	chOk := make(chan bool, 0)
// 	chOver := make(chan bool, 0)
// 	t := time.NewTimer(timeout)

// 	isTimeout := false
// 	go func(w *sync.WaitGroup, flag *bool) {
// 		w.Wait()
// 		if !isTimeout {
// 			chOk <- true
// 		} else {
// 			chOver <- true
// 		}
// 		fmt.Printf("ooooooooooover!!!\n")
// 	}(wg, &isTimeout)

// 	select {
// 	case <-chOk:
// 		// return isTimeout
// 		isTimeout = false
// 		break
// 	case <-t.C:
// 		isTimeout = true
// 		// 这里直接返回，会导致上面的goroutine泄露
// 	}

// 	mw := &sync.WaitGroup{}
// 	mw.Add(1)
// 	go func(w *sync.WaitGroup, m *sync.WaitGroup) {
// 		for {
// 			w.Done()
// 			select {
// 			case <-chOver:
// 				m.Done()
// 				return
// 			case <- time.After(time.Second):
// 				break
// 			}
// 		}
// 	}(wg, mw)

// 	mw.Wait()
// 	return isTimeout
// }

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false

	// 这里之所以采用容量为1的channel，是为了防止时候退出后，chDone没有读取，导致goroutine泄露
	chDone := make(chan bool, 1) 
	t := time.NewTimer(timeout)

	go func(w *sync.WaitGroup) {
		w.Wait()
		chDone <- false // 
	}(wg)

	isTimeout := false
	select {
	case <-chDone:
		return isTimeout
	case <-t.C:
		isTimeout = true
		return isTimeout
	}
}

