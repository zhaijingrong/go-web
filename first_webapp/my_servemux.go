/*
 * @Descripttion:
 * @version:
 * @Author: zjr
 * @Date: 2019-08-21 13:38:15
 * @LastEditors: zjr
 * @LastEditTime: 2019-08-21 13:48:04
 */
package main

import (
	"fmt"
	"net/http"
)

// 自定义一个多路复用器的结构
type MyMux struct{}

// 实现ServeHTTP方法
func (m MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 判断URL并转到对应的handler处理
	if r.URL.Path == "/index" {
		IndexHandler(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is index page")
}

func main() {
	// 实例化一个自定义的路由器
	mux := MyMux{}
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
