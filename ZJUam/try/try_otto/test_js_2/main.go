package main

import (
	"io/ioutil"
	"strings"

	"github.com/robertkrimen/otto"
)

func JsParser(filePath string, functionName string, args ...interface{}) (result string) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}
	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		panic(err)
	}

	return value.String()
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

func main() {
	rt := JsParser("./security.js", "")
}
