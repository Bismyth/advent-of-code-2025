package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) {
	numbers := [][]int{}
	answer := 0
	for i, r := range lines {

		if i < len(lines)-1 {
			j := 0
			for _, n := range strings.Split(r, " ") {
				if n == "" {
					continue
				}
				nR, _ := strconv.Atoi(n)

				if i == 0 {
					numbers = append(numbers, []int{nR})
				} else {
					numbers[j] = append(numbers[j], nR)
				}
				j++
			}

		} else {
			j := 0
			for _, s := range strings.Split(r, " ") {
				if s == "" {
					continue
				}

				if s == "+" {
					for _, n := range numbers[j] {
						answer += n
					}
				}
				if s == "*" {
					add := 1
					for _, n := range numbers[j] {
						add *= n
					}
					answer += add
				}
				j++
			}
		}

	}
	fmt.Println(answer)
}

func part2(lines []string) {

	offsets := []int{}

	currentCol := 0
	for _, c := range lines[len(lines)-1][1:] {
		if c == ' ' {
			currentCol += 1
		} else {
			offsets = append(offsets, currentCol)
			currentCol = 0
		}
	}
	offsets = append(offsets, currentCol+1)
	nums := [][]string{}
	for _, off := range offsets {
		nums = append(nums, make([]string, off))
	}
	answer := 0
	for rCount, r := range lines {

		if rCount != len(lines)-1 {

			currentIndex := 0
			for offIndex, off := range offsets {
				for i := 0; i < off; i++ {
					nums[offIndex][i] += string(r[currentIndex+i])
				}
				currentIndex += off + 1
			}
		} else {
			currentIndex := 0
			for offIndex, off := range offsets {

				s := r[currentIndex]
				actualNums := []int{}
				for _, n := range nums[offIndex] {
					nR, _ := strconv.Atoi(strings.Trim(n, " "))
					actualNums = append(actualNums, nR)
				}

				if s == '+' {
					for _, n := range actualNums {
						answer += n
					}
				}
				if s == '*' {
					add := 1
					for _, n := range actualNums {
						add *= n
					}
					answer += add
				}

				currentIndex += off + 1
			}
		}

	}
	fmt.Println(answer)

}

func main() {
	dat, err := os.ReadFile("./day-06-input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	part2(lines)

}
