package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// request body只能读一次，第二次读出来为空
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("first time read:\n%s\n", b)

		b1, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("second time read(should be blank):\n%s\n", b1)
	})

	http.HandleFunc("/requestdump", func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(requestDump))
	})

	http.ListenAndServe(":8000", nil)

}
