package main

import (
	"fmt"
)

func main() {
	var result int = 0
	for i := 1; i <= 100; i++ {
		result += i
	}
	fmt.Println(result)
}
