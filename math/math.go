package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Printf("%.4f\n", math.Cos(15.5))

	fmt.Printf("%.4f\n", math.Pow(3.5, 2))

	fmt.Printf("%.4f\n", math.Pow(math.Pi, 2)/6)

	c := math.Mod(7, 4)
	fmt.Printf("%.1f", c)
}
