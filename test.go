package main

import "fmt"

func main() {
	var x float64 = 32.8
	var y int = int(x)
	println(y)
	var z float64 = float64(y)
	fmt.Printf("%.2f\n", z)

}
