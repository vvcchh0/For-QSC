package main

import (
	_ "container/list"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/robertkrimen/otto"
)

func Encrypt(pwd string, exponent string, modulus string) string {
	var security_js_path = "./security.js"
	var reversedPwd string
	var result_ottoValue otto.Value
	var result string
	// 0. 将密码作倒序处理
	reversedPwd = reverse_string(pwd)
	// var bytes, err 用于记录 security.js 读取得到的数据
	//  1. 读取 js 内容
	fmt.Println("OK 1")
	bytes, err := ioutil.ReadFile(security_js_path)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK 2")
	//  var vm JS运行的引擎
	//  2. 创建引擎 vm

	vm := otto.New()
	fmt.Println("OK 3")
	// var useless 占位, VSCode报错看着难受
	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}
	fmt.Println("OK 4")
	//  var value 代码运行结果
	//  3. 运行 security.js 中的加密核心函数之一: getKeyPair, key 记录密钥
	key, err := vm.Call("RSAKeyPair", nil, exponent, "", modulus)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK 5")
	result_ottoValue, err = vm.Call("RSAUtils.encryptedString", nil, key, reversedPwd)
	if err != nil {
		panic(err)
	}
	result = result_ottoValue.String()
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
