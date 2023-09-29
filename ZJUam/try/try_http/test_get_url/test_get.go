// https://blog.csdn.net/qq_43778308/article/details/115448697 关于 resp.Body
// https://bbs.huaweicloud.com/blogs/c0473670c5e44f6d8aa411976e826de7  Get
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://zjuam.zju.edu.cn/cas/v2/getPubKey"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
