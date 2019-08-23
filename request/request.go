/*
 * @Descripttion:
 * @version:
 * @Author: zjr
 * @Date: 2019-08-22 10:56:40
 * @LastEditors: zjr
 * @LastEditTime: 2019-08-22 14:51:18
 */
package main

import (
	"fmt"
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	// Request中的URL结构
	fmt.Fprintln(w, r.URL.Path)
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	// Request结构中的Header是映射（map）类型，可以通过键值对进行操作，也可以使用map实现的Get, Set等方法操作。
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	// Request结构中的body字段实现了Reader接口，所以可以使用Read方法。
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	http.HandleFunc("/url", urlHandler)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/body", bodyHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
