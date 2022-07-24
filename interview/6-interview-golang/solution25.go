package main

// # 高并发下的锁与map的读写

// 场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。

// 每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	vlock    sync.RWMutex
}

func NewBan() *Ban {

	b := &Ban{visitIPs: make(map[string]time.Time)}

	return b
}

func (b *Ban) loop(ctx context.Context) {
	duration := time.Minute * 3
	go func() {
		t := time.NewTimer(duration)
		for {
			select {
			case <-t.C:
				b.vlock.Lock()
				for k, t := range b.visitIPs {
					now := time.Now()
					if now.Sub(t) > duration {
						delete(b.visitIPs, k)
					}
				}
				t.Reset(duration)
				b.vlock.Unlock()
			case <-ctx.Done(): // 结束当前goroutine
				return
			}
		}
	}()
}

func (b *Ban) isBan(ip string) bool {
	b.vlock.Lock()
	defer b.vlock.Unlock()
	if _, ok := b.visitIPs[ip]; ok {
		return true
	}
	return false
}

func (b *Ban) insertBan(ip string) {
	b.vlock.Lock()
	defer b.vlock.Unlock()
	b.visitIPs[ip] = time.Now()
}

func (o *Ban) visit(ip string) bool {
	if o.isBan(ip) {
		return false
	}
	o.insertBan(ip)
	return true
}

func main() {
	success := 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 主进程结束前 ctx.Done()会收到结束信号

	ban := NewBan()
	ban.loop(ctx)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(n int) {
				ip := fmt.Sprintf("192.168.1.%d", n)
				if ban.visit(ip) {
					success++
				}
			}(j)
		}
	}
	fmt.Println("success:", success)
}
