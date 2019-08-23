/*
 * @Descripttion:
 * @version:
 * @Author: zjr
 * @Date: 2019-08-23 10:32:30
 * @LastEditors: zjr
 * @LastEditTime: 2019-08-23 10:55:14
 */
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//
		r.ParseMultipartForm(1024)
		// 使用r.FormFile获取文件句柄，然后对文件进行存储等处理。
		file, handler, err := r.FormFile("uploadfile")
		defer file.Close()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprintf(w, "%v", handler.Header)

		// 假设已经有upload目录，存储文件。
		f, err := os.OpenFile("../static/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	} else {
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, nil)
	}

}

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
