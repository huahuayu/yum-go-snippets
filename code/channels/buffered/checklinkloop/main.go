package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://qq.com",
		"http://taobao.com",
		"http://baidu.com",
		"http://z.cn",
		"http://douban.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) { // could be any routine
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}