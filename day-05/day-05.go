package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./day-05-input.txt")
	if err != nil {
		panic(err)
	}

	ranges := [][2]int{}

	data := strings.Split(string(dat), "\n\n")

	for _, r := range strings.Split(data[0], "\n") {
		nums := strings.Split(r, "-")

		if len(nums) != 2 {
			panic("bad numbers")
		}

		d1, _ := strconv.Atoi(nums[0])
		d2, _ := strconv.Atoi(nums[1])

		ranges = append(ranges, [2]int{d1, d2})
	}

	// answer := 0

	// for _, d := range strings.Split(data[1], "\n") {
	// 	num, _ := strconv.Atoi(d)

	// 	for _, r := range ranges {
	// 		if num >= r[0] && num <= r[1] {
	// 			answer += 1
	// 			break
	// 		}
	// 	}
	// }

	slices.SortFunc(ranges, func(a [2]int, b [2]int) int {
		return a[0] - b[0]
	})

	combined := [][2]int{ranges[0]}
	for _, r := range ranges[1:] {

		lastEntry := len(combined) - 1

		if r[0] <= combined[lastEntry][1] {
			if r[1] > combined[lastEntry][1] {
				combined[lastEntry][1] = r[1]
			}
		} else {
			combined = append(combined, r)
		}
	}
	answer := 0
	for _, r := range combined {
		answer += r[1] - r[0] + 1
	}
	fmt.Println(answer)
}
