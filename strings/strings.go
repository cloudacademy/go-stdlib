package main

import (
	"fmt"
	"strings"
)

func main(){
	//https://pkg.go.dev/strings

	var1 := "cloud"
	var2 := "academy"
	var3 := var1 + var2
	upper := strings.ToUpper(var3)
	hascloud := strings.HasPrefix(var3, var1)
	idx := strings.Index(var3, var2)
	strs := strings.Split(var3, "a")
	aca := strings.Join(strs, "**")

	fmt.Println(upper)
	fmt.Println(hascloud)
	fmt.Println(idx)
	fmt.Println(aca)

	convert := func(r rune) rune {
		if r == 'a' {
			return '@'
		}
		if r == 'o' {
			return '0'
		}
		if r == 'e' {
			return '3'
		}

		return r
	}

	fmt.Println(strings.Map(convert, var3))
}