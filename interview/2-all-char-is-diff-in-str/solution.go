package main

import "fmt"

/*
## 判断字符串中字符是否全都不同

**问题描述**

请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。

输入保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

*/

func solution(str string) bool {
	for i := 0; i < len(str) - 1; i++ {
		for j := i+1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}

func main() {

	tests := make(map[string]bool, 0)
	tests["abc"] = true
	tests["a12345678a"] = false
	tests["abcdefghijk"] = true
	tests["a"] = true
	tests["aa"] = false
	tests[""] = true

	for k,v := range tests {
		if v != solution(k) {
			fmt.Printf("test case: %s, expected: %v, got: %v\n", k, v, solution(k))
		}
	}
}
