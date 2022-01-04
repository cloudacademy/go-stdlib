package main

import (
	"fmt"
	"os"
	"os/exec"
	"bytes"
)

func main(){
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