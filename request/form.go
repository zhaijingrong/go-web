/*
 * @Descripttion:
 * @version:
 * @Author: zjr
 * @Date: 2019-08-22 14:27:42
 * @LastEditors: zjr
 * @LastEditTime: 2019-08-22 14:56:57
 */
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	// 登录逻辑，GET方法时放回登录页面，POST解析用户、密码进行登录验证。
	r.ParseForm() // 解析url传递的参数，对于POST请求则解析请求体（request body）
	if r.Method == "POST" {
		fmt.Println(r.Form["username"])
		fmt.Println(r.Form["password"])
	} else {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
