package main

import (
	"fmt"
	"math"
	"strconv"
)

func convertbin2dec(binNum string) {
	if !validateInput(binNum) {
		return
	}

	decNum := 0

	for i, num := range binNum {
		numInt, _ := strconv.Atoi(string(num))
		decNum += int(math.Pow(float64(numInt)*2, float64(i)))
	}

	fmt.Println(binNum, "in binary equals to", decNum, "in decimal")
}

func validateInput(number string) bool {
	if numIncorrectLength(number) || !inputIsOnlyOnesAndOrZeros(number) {
		fmt.Printf("incorrect number")
		return false
	}
	return true
}

func numIncorrectLength(n string) bool {
	return len(n) > 8
}

func inputIsOnlyOnesAndOrZeros(n string) bool {
	for _, num := range n {
		if num < 48 || num > 49 {
			return false
		}
	}
	return true
}
