package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func largestNumber(s string) int {

	largest := 0
	second := 0

	for _, c := range s[:len(s)-1] {
		v := int(c) - 48

		if v > largest {
			largest = v
			second = 0
			continue
		}

		if v > second {
			second = v
		}
	}
	final := int(s[len(s)-1]) - 48
	if final > second {
		second = final
	}
	res := largest*10 + second

	return res
}

func largestNumberWithDigits(s string, digits int) int {

	strLength := len(s)
	currentStart := 0
	ans := 0
	for x := 0; x < digits; x++ {
		endSize := digits - 1 - x
		currentNum := 0
		for i := currentStart; i < strLength-endSize; i++ {
			v := int(s[i]) - 48
			if v > currentNum {
				currentNum = v
				currentStart = i + 1
			}
		}

		ans += currentNum * int(math.Pow(10, float64(digits-x-1)))
	}

	return ans
}

func main() {
	dat, err := os.ReadFile("./day-03-test.txt")
	if err != nil {
		panic(err)
	}

	answer := 0

	for _, l := range strings.Split(string(dat), "\n") {
		answer += largestNumberWithDigits(l, 12)
	}

	fmt.Println(answer)
}
