package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func init() {
	fmt.Println("init method called")
	// 检查1：os.Args[0] = 程序名 ; os.Args[1] = 第一个参数名; 必须是两个参数
	if len(os.Args) != 2 {
		fmt.Println("Usage: gocurl <url>")
		os.Exit(1)
	}

	// 必须http或https开头
	if !(strings.HasPrefix(os.Args[1],"http://") || strings.HasPrefix(os.Args[1],"https://")){
		fmt.Println("url should be prefix with http:// or https://")
		os.Exit(1)
	}
}

// 调用系统curl命令
func invokecurl(url string){
	cmd := exec.Command("curl",url)
	output, err := cmd.CombinedOutput()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(string(output))
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(gohtml.Format(string(body)))
	fmt.Println("_____________ gocurl command output______________")
	fmt.Println(string(body))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("_____________system curl command output______________")
	invokecurl(os.Args[1])
}
