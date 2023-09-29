package main

import (
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

func JsParser(filePath string, functionName string, args ...interface{}) (result string) {
	bytes, err := ioutil.ReadFile(filePath)
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

func main() {
	//rt := JsParser("./test.js", "add", 1, 2)
	rt := JsParser("./test2.js", "add", 1, 2)
	fmt.Println(rt)
}
