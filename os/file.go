package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	bytes, err := os.ReadFile("makefile")
	if err != nil {
		log.Fatal(err)
	}

	data := string(bytes)

	fmt.Println(data)

	err = os.WriteFile("somenewfile", bytes, 0666)
	if err != nil {
		log.Fatal(err)
	}
}