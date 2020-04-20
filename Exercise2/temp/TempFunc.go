package temp

func dividing(beDivNum int, divNum float64) float64 {
	var answer float64
	answer = float64(beDivNum) / divNum
	return answer
}

func DivideCall(beDivNum int, divNum float64) float64 {
	return dividing(beDivNum, divNum)
}
