package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

func main(){
	text := `Learning Objectives
By completing this course, you will:

Learn about what makes Go a great language
Learn how to install the Go toolchain
Learn how to setup Visual Studio Code to edit and debug Go programs
Learn how to work with the Go Playground to test and run snippets of Go code
Learn and understand the basic Go language syntax and features
Learn how to use the Go tool chain commands to compile, test, and manage Go code
And finally, you’ll learn how to work with and manage Go modules for module dependency management`

	//owner can change it, everyone else can read it.
	//view permissions: stat -f %A ./io/sampleout.txt
	var filePerm fs.FileMode = 0644

	f, err := os.OpenFile(os.Args[1], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := strings.NewReader(text)
	buffer := make([]byte, 10)
	for {
		n, err := s.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			f.Write(buffer[:n])
		}
	}	
}