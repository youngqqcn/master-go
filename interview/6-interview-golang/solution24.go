// # 实现阻塞读且并发安全的map

// GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：

// ```go
// type sp interface {
//     Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
//     Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
// }
// ```
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type sp interface {
	//存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Write(key string, val interface{})

	//读取一个key，如果key不存在阻塞，等待key存在或者超时
	Read(key string, timeout time.Duration) interface{}
}

type MyMap struct {
	data   map[string]interface{}
	_q     map[string]chan interface{} // 将等待的groutine的chan注册
	_qlock sync.RWMutex
	_dlock sync.RWMutex
}

func NewMyMap() *MyMap {
	return &MyMap{
		data: make(map[string]interface{}, 0),
		_q:   make(map[string]chan interface{}, 0),
	}
}

func (m *MyMap) register(key string) chan interface{} {
	m._qlock.Lock()
	defer m._qlock.Unlock()
	ch := make(chan interface{}, 0)
	m._q[key] = ch
	return ch
}

func (m *MyMap) unregister(key string) {
	m._qlock.Lock()
	defer m._qlock.Unlock()
	close(m._q[key])
	delete(m._q, key)
}

func (m *MyMap) Write(key string, val interface{}) {
	go func(x *MyMap, v interface{}) {
		m._qlock.Lock()
		defer m._qlock.Unlock()
		if ch, ok := m._q[key]; ok {
			ch <- v
		}
	}(m, val)

	m._dlock.Lock()
	defer m._dlock.Unlock()
	m.data[key] = val
}

func (m *MyMap) Read(key string, timeout time.Duration) interface{} {

	m._dlock.Lock()
	if v, ok := m.data[key]; ok {
		m._dlock.Unlock()
		return v
	}
	m._dlock.Unlock()

	ch := m.register(key)
	defer m.unregister(key)

	for {
		select {
		case v := <-ch:
			return v
		case <-time.After(timeout):
			return nil
		}
	}
}

func main() {
	mp := NewMyMap()

	wg := &sync.WaitGroup{}

	wg.Add(20 + 20)

	var counter int64 = 0

	// read goroutines
	for i := 0; i < 20; i++ {
		go func(n int, m *MyMap, cnt *int64) {
			for j := 0; j < 10000; j++ {
				_ = mp.Read(fmt.Sprintf("%d-%04d", n, j), time.Second*30)
				atomic.AddInt64(cnt, 1)
			}
			wg.Done()
		}(i, mp,  &counter)
	}

	// write goroutines
	for i := 0; i < 20; i++ {
		go func(n int) {
			for j := 0; j < 10000; j++ {
				mp.Write(fmt.Sprintf("%d-%04d", n, j), j)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
}
