package main

import (
	"fmt"
	"net/url"
)

func main() {
	var url_string string
	url_string = "https://zjuam.zju.edu.cn/cas/v2/getPubKey"
	str_url, err := url.Parse(url_string)
	if err != nil {
		panic(err)
	}
	queryValues, _ := url.ParseQuery(str_url.RawQuery)
	fmt.Println("Scheme: ", str_url.Scheme)
	fmt.Println("Host: ", str_url.Host)
	fmt.Println("Port: ", str_url.Port())
	fmt.Println("Path: ", str_url.Path)
	fmt.Println("RawQuery: ", str_url.RawQuery)
	fmt.Println("Fragment: ", str_url.Fragment)
	fmt.Println("modulus: ", queryValues.Get("modulus"))
	fmt.Println(queryValues)
}
