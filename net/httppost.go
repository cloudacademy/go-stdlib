package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const HTTPBIN_URL = "https://httpbin.org/post?course=golang&author=jeremycook"

func PostRequest1() {
	data := `{"fruit": "Apple","size": "Large","color": "Red"}`
	bodyreader := strings.NewReader(data)

	resp, err := http.Post(HTTPBIN_URL, "application/json", bodyreader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Printf(string(body))
}

func PostRequest2() {
	data := `{"fruit": "Apple","size": "Large","color": "Red"}`
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var jsonStr = []byte(data)
	req, _ := http.NewRequest("POST", HTTPBIN_URL, bytes.NewBuffer(jsonStr))
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
	PostRequest1()
	fmt.Println()
	PostRequest2()
}
