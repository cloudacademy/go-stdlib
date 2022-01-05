package main

import (
	"fmt"
	"os"
)

func main() {
	//https://pkg.go.dev/os

	port := os.Getenv("CLOUDACADEMY_PORT")
	host := os.Getenv("CLOUDACADEMY_HOST")

	fmt.Println(port)
	fmt.Println(host)
}
