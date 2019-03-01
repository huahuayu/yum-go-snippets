package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的, 如果不解析，则下面r.Form是零值
	//以下是输出到服务器端的打印信息
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("r.Form[\"firstname\"]", r.Form["fistname"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	//这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "Hello shiming!")
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
