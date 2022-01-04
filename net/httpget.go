package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const HTTPBIN_URL = "https://httpbin.org/get?course=golang&author=jeremycook"

func GetRequest1() {
	resp, err := http.Get(HTTPBIN_URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Printf(string(body))
}

func GetRequest2() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", HTTPBIN_URL, nil)
	req.Header.Set("Accept", "*/*")
	resp, err := client.Do(req)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		// timeout error
		panic("TIMEOUT!!")
	} else if err != nil {
		// error, but not a timeout
		panic("ERROR!!")
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Printf(string(body))
}

func main() {
	GetRequest1()
	GetRequest2()
}
