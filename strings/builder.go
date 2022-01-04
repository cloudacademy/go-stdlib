package main

import (
	"fmt"
	"strings"
)

func main() {
	//https://pkg.go.dev/strings

	var b strings.Builder
	for i := 1; i <= 10; i++ {
		//b.WriteString(fmt.Sprintf("%d......", i*i))
		fmt.Fprintf(&b, "%d....", i*i*i)
	}
	b.WriteString("done")
	fmt.Println(b.String())

}