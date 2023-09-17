package main

import (
	"fmt"

	"strconv"
)

func main() {
	var a string = "2"
	var b int = 3
	c, err := strconv.Atoi(a)

	if err != nil {
		panic(err)
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		fmt.Println(c + b)
	}
}
