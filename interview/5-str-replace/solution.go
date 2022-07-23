package main

import (
	"fmt"
	"strings"
)

/*
## 字符串替换问题

**问题描述**

请编写一个方法，将字符串中的空格全部替换为“%20”。
假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。
给定一个string为原始的串，返回替换后的string。

*/

func solution(s string) string {
	s = strings.ReplaceAll(s, " ", "%20")
	return s
}

func main() {

	tests := make(map[string]string, 0)
	tests["a b c d"] = "a%20b%20c%20d"

	for k, v := range tests {
		if ret := solution(k); ret != v {
			fmt.Printf("failed, expected: %v, got: %v\n", v, ret)
		}
	}

}
