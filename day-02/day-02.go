package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRepeated(n int) bool {

	sample := strconv.Itoa(n)
	if len(sample) == 1 {
		return false
	}
	if len(sample) == 2 {
		return sample[0] == sample[1]
	}

	for x := 1; x <= len(sample)/2; x += 1 {
		if len(sample)%x != 0 {
			continue
		}
		repeats := len(sample) / x
		substr := sample[0:x]

		if strings.Repeat(substr, repeats) == sample {
			return true
		}
	}

	return false
}

func isRepeatedOnce(n int) bool {
	sample := strconv.Itoa(n)
	if len(sample)%2 != 0 {
		return false
	}

	cutSize := len(sample) / 2
	return strings.Repeat(sample[0:cutSize], 2) == sample
}

func main() {
	dat, err := os.ReadFile("./day-02-input.txt")
	if err != nil {
		panic(err)
	}

	ranges := [][]int{}

	for _, r := range strings.Split(string(dat), ",") {
		rArr := []int{}
		for _, n := range strings.Split(r, "-") {
			parsed, err := strconv.Atoi(strings.Trim(n, " \n"))
			if err != nil {
				fmt.Printf("bad number: %s\n", n)
				return
			}
			rArr = append(rArr, parsed)
		}
		ranges = append(ranges, rArr)
	}
	answer := 0
	for _, r := range ranges {
		for x := r[0]; x <= r[1]; x++ {
			if isRepeated(x) {
				answer += x
			}
		}
	}
	fmt.Println(answer)
}
