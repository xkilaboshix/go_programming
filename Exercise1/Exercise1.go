package main

import (
	"fmt"
)

func dividing(beDivNum int, divNum float64) float64 {
	var answer float64
	answer = float64(beDivNum) / divNum
	return answer
}

func main() {
	var beDivNum int = 20
	var divNum float64 = 0.002
	fmt.Println(dividing(beDivNum, divNum))
}
