/*
向网站发送数据
参考:  https://www.bilibili.com/video/BV1XJ411v79e
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postForm(username string, password string, execution string) {
	//var url_string string = "https://zjuam.zju.edu.cn/cas/login"
	var url_string string = "https://identity.zju.edu.cn/auth/realms/zju/broker/cas-client/endpoint?state=aMHihcKA5c3k-WHv2-TinqGnG_8-rd95399-F0sLoV8.Ihkc9DtxTfo.TronClass"
	var contentType string = "application/x-www-form-urlencoded"

	//  构建 Form : data
	data := make(url.Values)
	data.Add("username", username)
	data.Add("password", password)
	data.Add("authcode", "")
	data.Add("execution", execution)
	data.Add("_eventId", "submit")
	payload := data.Encode()

	r, _ := http.Post(url_string, contentType, strings.NewReader(payload))
	defer func() { _ = r.Body.Close() }()

	content, _ := ioutil.ReadAll(r.Body) // 读取提交Form后网页返回的信息
	fmt.Println(string(content))
}
