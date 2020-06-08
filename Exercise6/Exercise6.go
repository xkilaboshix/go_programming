package main

import (
	"fmt"
)

func main() {
	var slice []int

	for i := 1; i <= 10; i += 2 {
		slice = append(slice, i)
	}
	fmt.Println(slice)
}
