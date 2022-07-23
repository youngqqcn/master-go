package main

import "fmt"

/*
## 翻转字符串

**问题描述**

请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
*/

func solution(s string) (string, bool){
	
	if len(s) > 5000 {
		return s, false
	}

	// b := []byte(s)
	b := []rune(s)

	i := 0
	j := len(b) - 1
	for ; i < j;  {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b), true
}

func main() {
	tests := make(map[string]string, 0)
	tests["abcdefgh"] = "hgfedcba"
	tests["abc"] = "cba"
	tests["123456789"] = "987654321"
	tests["1"] = "1"
	tests["我爱中国"] = "国中爱我"

	for k, v := range tests {
		if ret, ok := solution(k) ; !ok{
			fmt.Printf("expected: %v, got: %v\n", v, ret)
		}
	}
}
