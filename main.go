package main

import "fmt"

func main() {
	array := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	lol := [5]int{0, 1, 2, 3, 4}

	fmt.Println(array)
	fmt.Println(lol[0], len(lol)-1)
}
