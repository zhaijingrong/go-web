/*
 * @Descripttion:
 * @version:
 * @Author: zjr
 * @Date: 2019-08-20 09:56:16
 * @LastEditors: zjr
 * @LastEditTime: 2019-08-20 10:09:41
 */
package main

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	//将“Hello World"字符串写入到ResponseWriter, 并格式化字符串，将当前请求的路径添加进去
	fmt.Fprintf(w, "Hello world, %s!", r.URL.Path[1:])
}

func main() {
	//将定义好的HelloWorldHandler函数设置成根（/）URL被访问时的处理器
	http.HandleFunc("/", HelloWorldHandler)

	//启动服务器并监听8000/tcp端口
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
