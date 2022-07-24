package main

import (
	"fmt"
	"sync"
)

// 下面的代码有什么问题?

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

// func (ua *UserAges) Get(name string) int {
// 	if age, ok := ua.ages[name]; ok {
// 		return age
// 	}
// 	return -1
// }

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

// test
func main() {
	ua := &UserAges{ages: make(map[string]int, 0)}

	wg := &sync.WaitGroup{}
	wg.Add(10 + 4)

	// write
	for gid := 10000; gid < 10010; gid++ {
		go func(gid int, u *UserAges) {
			for i := 0; i < 100000; i++ {
				name := fmt.Sprintf("gid%d-%d", gid, i)
				u.Add(name, i)
			}
			wg.Done()
		}(gid, ua)
	}

	// read
	for gid := 10000; gid < 10004; gid++ {
		go func(gid int, u *UserAges) {
			for i := 0; i < 100000; i++ {
				name := fmt.Sprintf("gid%d-%d", gid, i)
				_ = u.Get(name)
			}
			wg.Done()
		}(gid, ua)
	}

	wg.Wait()
}

// 但是map是并发读写不安全的。map属于引用类型，
// 并发读写时多个协程见是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系。会报错误信息:
// fatal error: concurrent map read and map write


// sync.Mutex 和 sync.RWMutex
//
// Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
//
// RWMutex 为读写锁，
//     写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占
//     适用于读多写少的场景
