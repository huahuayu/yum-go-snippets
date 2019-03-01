package httpbody

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://baidu.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//lw := logWriter{}

	//io.Copy(lw, resp.Body)
	io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
