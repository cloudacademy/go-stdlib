package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//https://pkg.go.dev/os

	fmt.Println(os.Getenv("PATH"))

	var outb, errb bytes.Buffer
	cmd := exec.Command("helm")
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(outb.String())
}
