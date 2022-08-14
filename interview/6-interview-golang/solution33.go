package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu  // 锁的状态也会被复制
	mu.count++
	mu.Unlock()
	// mu2.Lock()   // mu2已经是lock状态
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}
