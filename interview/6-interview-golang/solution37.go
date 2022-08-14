package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	m.Store("xxxx", "ffff")
	m.Store(444, 222)

	fmt.Println(m.Load(444))
}
