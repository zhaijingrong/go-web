/*
 * @Descripttion:
 * @version:
 * @Author: zjr
 * @Date: 2019-08-21 11:37:27
 * @LastEditors: zjr
 * @LastEditTime: 2019-08-22 14:51:19
 */
package main

import (
	"fmt"
	"net/http"
)

// 定一个自定义的Handler的结构
type MyHanlder struct{}

// 为MyHanlder结构实现Hanlder接口的ServeHTTP的方法，此时MyHandler将是一个处理器(Handler)
func (h MyHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This my handler")
}

func main() {
	// 实例化MyHandler
	hanlder := MyHanlder{}

	//启动服务器并监听8000/tcp端口
	err := http.ListenAndServe(":8000", hanlder)
	if err != nil {
		fmt.Println(err)
	}
}
