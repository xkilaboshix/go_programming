package main

import (
	"fmt"
)

func main() {
	var result [10]int

	for i := 1; i <= 10; i++ {
		result[i-1] = i
	}
	fmt.Println(result)

}
