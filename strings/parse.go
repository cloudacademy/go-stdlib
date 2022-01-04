package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	//https://pkg.go.dev/strconv

	var1 := 100
	var2 := math.Pi
	var3 := "200"

	var4 := strconv.Itoa(var1)
	var5, err := strconv.Atoi(var3)
	if err != nil {
		panic("whoops...")
	}

	fmt.Printf("%d, %f, %s, %s, %d\n", var1, var2, var3, var4, var5)
	fmt.Printf("%T, %T, %T, %T, %T\n", var1, var2, var3, var4, var5)

	f := "3.14159265"
	var7, _ := strconv.ParseFloat(f, 64)
	fmt.Printf("%f -- %T\n", var7, var7)

	s := fmt.Sprintf("%.4f", var7)
	fmt.Println(s)
}