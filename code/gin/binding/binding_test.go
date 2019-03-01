package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestBinding(t *testing.T) {
	url := "http://localhost:8080/loginJSON"

	payload := strings.NewReader("{\n\t\"user\":\"manu\",\n\t\"password\":123,\n\t\"weight\":456\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "a9bcac88-ee30-4632-bce8-0214b506a0a5")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))
}
