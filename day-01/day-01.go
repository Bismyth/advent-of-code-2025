package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {

	dial := 50
	answer := 0
	for _, l := range lines {

		d := l[0]
		n, err := strconv.Atoi(l[1:])
		if err != nil {
			fmt.Printf("bad number: %s\n", l[1:])
			return -1
		}
		n = n % 100
		if d == 'L' {
			dial -= n
			if dial < 0 {
				dial += 100
			}
		}
		if d == 'R' {
			dial += n
			if dial > 99 {
				dial -= 100
			}
		}
		if dial == 0 {
			answer += 1
		}
	}

	return answer
}

func part2(lines []string) int {
	dial := 50
	answer := 0
	for _, l := range lines {

		d := l[0]
		n, err := strconv.Atoi(l[1:])
		if err != nil {
			fmt.Printf("bad number: %s\n", l[1:])
			return -1
		}

		answer += n / 100
		n = n % 100

		on0 := dial == 0

		if d == 'L' {
			dial -= n
			if dial < 0 {
				dial += 100
				if !on0 {
					answer += 1
				}
			}
		}
		is100 := false
		if d == 'R' {
			dial += n
			if dial == 100 {
				is100 = true
			}
			if dial > 99 {
				dial -= 100
				if !on0 {
					answer += 1
				}
			}
		}

		if dial == 0 && !is100 {
			answer += 1
		}
	}

	return answer
}

func main() {

	dat, err := os.ReadFile("./day-01-input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	fmt.Println(part2(lines))

}
