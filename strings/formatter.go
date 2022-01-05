package main

import (
	"fmt"
	"math"
)

type Rocket struct {
	Name   string
	Engine string
	Thrust float32
	Length int
}

func main() {
	//https://pkg.go.dev/fmt

	var1 := 3
	var2 := "cloudacademy"
	var3 := math.Pi
	var4 := var1 > 2

	fmt.Printf("blah")

	fmt.Printf("integer -> %d\n", var1)
	fmt.Printf("string -> %s\n", var2)
	fmt.Printf("float -> %f\n", var3)
	fmt.Printf("float -> %.2f\n", var3)
	fmt.Printf("float -> %.4f\n", var3)
	fmt.Printf("bool -> %t\n", var4)

	arr := []interface{}{var1, var2, var3}

	fmt.Printf("item3 %[3]f, item2 %[2]s, item1 %[1]d\n", arr[0], arr[1], arr[2])
	fmt.Printf("item3 %[3]T, item2 %[2]T, item1 %[1]T\n", arr[0], arr[1], arr[2])

	saturnV := Rocket{
		Name:   "saturnV",
		Engine: "F-1",
		Thrust: 1.5,
		Length: 42,
	}

	fmt.Printf("%v\n", saturnV)
	fmt.Printf("%+v\n", saturnV)
	fmt.Printf("%T\n", saturnV)

	var5 := fmt.Sprintf("area of circle radius 4 = %f", 2*var3*4)
	fmt.Println(var5)
}
