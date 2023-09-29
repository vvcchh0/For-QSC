/*
加密相关代码
从 security.js 直译

关于JS中部分代码的说明:
String.charCodeAt(index) 效果同 GO 中的 String[i]
String.substr(a, b) 效果同 GO 中的 string(String[a:a+b])
*/
package main

import (
	"container/list"
	_ "fmt"
	_ "strings"
)

func biCopy(bi list.List) list.List {
	var result list.List
	// var i用于正向遍历

	for i := bi.Front(); i != nil; i = i.Next() {
		result.PushBack(i.Value)
	}

	return result
}
func charToHex(bt byte) int { // JS: RSAUtils.charToHex Line 169
	var ZERO byte = byte('0')
	var NINE byte = byte('9')
	var a byte = byte('a')
	var f byte = byte('f')
	var A byte = byte('A')
	var F byte = byte('F')
	var result int

	if ZERO <= bt && bt <= NINE {
		result = int(bt - ZERO)
	} else if a <= bt && bt <= f {
		result = 10 + int(bt-a)
	} else if A <= bt && bt <= F {
		result = 10 + int(bt-A) // 虽然但是, modulus中仅有数字与小写字母
	}

	return result
}

func hexToDigit(s string) int { // JS: RSAUtils.hexToDigit Line 190
	var result int = 0
	var s1 = min(len(s), 4)
	var i int
	for i = 0; i < s1; i++ {
		result <<= 4
		result |= charToHex(charCodeAt(s, i))
	}
	return result
}

func biFromHex(s string) list.List { // RSAUtils.biFromHex Line 200
	var result list.List
	var s1 = len(s)
	var i int

	for i = s1; i > 0; {
		result.PushBack(hexToDigit(substr(s, max(i-4, 0), min(i, 4))))
		i -= 4
	}

	return result
}

func biHighIndex(x list.List) int { // RSAUtils.biHighIndex Line 289
	var result int = x.Len() - 1
	// var i 用于反向遍历

	for i := x.Back(); i != nil && result > 0 && i.Value == 0; i = i.Prev() {
		result--
	}

	return result
}

func substr(s string, a int, b int) string { // JS: s.substr(a, b)
	var i int
	var data []byte
	var result string
	for i = a; i < a+b; i++ {
		data = append(data, s[i])
	}
	result = string(data[:])
	return result
}

func charCodeAt(s string, index int) byte { // JS: s.charCodeAt(i)
	return s[index]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
