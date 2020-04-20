package main

import (
	"fmt"

	"./temp"
)

func main() {
	var beDivNum int = 20
	var divNum float64 = 0.002
	fmt.Println(temp.DivideCall(beDivNum, divNum))
}
