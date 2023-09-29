package main

import (
	"container/list"
	"fmt"
	"strings"
)

func sayhello() {
	var l list.List
	var i int
	var s_array []string
	var data [10]byte
	data[0] = 'T'
	data[1] = 'E'
	var str string = string(data[:])
	i = 0
	fmt.Println("Hello!")
	fmt.Println(strings.Count("Hello!", "") - 1)
	fmt.Println(string("Hello!"[2]))
	for i < 6 {
		l.PushBack(string("Hello!"[i]))
		s_array = append(s_array, string("Hello!"[i]))
		i++
	}
	i = 0
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println(string(s_array[0]))
	fmt.Println(str)
	fmt.Println(string(byte('A')))
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

func reverse_string(s string) string {
	var reversed_string []string
	var length_string int
	var i int

	length_string = len(s)

	for i = length_string - 1; i >= 0; i-- { // 逆序遍历 字符串 s, 将取出的 byte 转化为 string 存入 字符串数组 reserved_string 中
		reversed_string = append(reversed_string, string(s[i]))
	}

	return strings.Join(reversed_string, "")
}

func test_two(a int, b int) (int, int) {
	return a, b
}
