package main

import (
	"fmt"
	"os"
)

func main() {
	//https://pkg.go.dev/os

	args1 := os.Args
	args2 := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(args1)
	fmt.Println(args2)
	fmt.Println(arg)
}
