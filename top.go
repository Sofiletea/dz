package main

import (
	"fmt"

	"strconv"
)

func main() {
	s := "97"
	if n, err := strconv.Atoi(s); err == nil {
		fmt.Println(n + 1)
	} else {
		fmt.Println(s)
	}
}
