package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

/*
## 判断两个给定的字符串排序后是否一致

**问题描述**

给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。
保证两串的长度都小于等于5000。

*/

func solution1(s1, s2 string) bool {
	if !(len(s1) <= 5000 && len(s2) <= 5000) {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	bz1, bz2 := []byte(s1), []byte(s2)
	sort.Slice(bz1, func(i, j int) bool {
		return bz1[i] < bz1[j]
	})
	sort.Slice(bz2, func(i, j int) bool {
		return bz2[i] < bz2[j]
	})

	if bytes.Compare(bz1, bz2) == 0 {
		return true
	}

	return false
}

func solution(s1, s2 string) bool {

	if len(s1) > 5000 || len(s2) > 5000 {
		return false
	}

	visited := make(map[byte]bool, 0)
	for i := 0; i < len(s1); i++ {
		if !visited[s1[i]] {
			ch := s1[i : i+1]
			visited[s1[i]] = true
			if strings.Count(s1, ch) != strings.Count(s2, ch) {
				return false
			}
		}
	}
	return true
}

func main() {

	tests := make([][2]string, 0)
	tests = append(tests, [2]string{"abcdefghi", "ihgfedcba"})
	tests = append(tests, [2]string{"1234567890", "0987612345"})
	tests = append(tests, [2]string{"fsjsddfsdf", "sfskfsdjfs"})
	expected := []bool{true, true, false}

	for i := 0; i < len(tests) && i < len(expected); i++ {
		test := tests[i]
		e := expected[i]
		if e != solution(test[0], test[1]) {
			fmt.Printf("testcase: %v , expected: %v, got: %v\n", i, e, !e)
		}

	}

}
