package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）, 如果不用ParseForm整体解析也可以通过r.FormValue按key进行解析
	log.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		cwd, _ := os.Getwd()
		//fmt.Println("filepath is:------------",filepath.Join( cwd, "./template/login.gtpl" ))
		t, err := template.ParseFiles(filepath.Join(cwd, "./template/login.gtpl"))
		if err != nil {
			log.Println("parse file error ", err)
		}

		err = t.Execute(w, nil)
		log.Println("parse template error: ", err)
	} else {
		userName := r.FormValue("username")
		password := r.FormValue("password")
		//请求的是登录数据，那么执行登录的逻辑判断
		log.Printf("[user login] username: %s, password: %s\n", userName, password)
		fmt.Fprintf(w, "your user name is %s, password is %s \n", userName, password)
	}
}

func main() {
	//http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
