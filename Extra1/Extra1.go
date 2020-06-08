package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var multipuls3 []string
	var multipuls5 []float64
	var multipuls15 []int

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			multipuls15 = append(multipuls15, i)
		} else if i%5 == 0 {
			multipuls5 = append(multipuls5, float64(i)/10)
		} else if i%3 == 0 {
			multipuls3 = append(multipuls3, strconv.Itoa(i)+"個")
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(multipuls15)))

	for i := 0; i < len(multipuls15); i++ {
		fmt.Println(multipuls15[i])
	}
	for i := 0; i < len(multipuls5); i++ {
		fmt.Println(multipuls5[i])
	}
	for i := 0; i < len(multipuls3); i++ {
		fmt.Println(multipuls3[i])
	}
}
