package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	x float64
	y float64
	z float64
}

func getDistance(p1, p2 Position) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2) + math.Pow(p1.z-p2.z, 2))
}

type Connection struct {
	distance   float64
	pointIndex int
}

func main() {

	dat, err := os.ReadFile("./day-08-input.txt")
	if err != nil {
		panic(err)
	}

	positions := []Position{}
	for _, l := range strings.Split(string(dat), "\n") {
		nums := []int{}
		for _, s := range strings.Split(l, ",") {
			i, _ := strconv.Atoi(s)
			nums = append(nums, i)
		}
		if len(nums) != 3 {
			panic("bad numbers")
		}

		p := Position{
			x: float64(nums[0]),
			y: float64(nums[1]),
			z: float64(nums[2]),
		}
		positions = append(positions, p)
	}

	relations := make([][]Connection, len(positions))

	for i, p1 := range positions {
		for jR, p2 := range positions[i+1:] {
			j := i + jR + 1
			distance := getDistance(p1, p2)
			relations[i] = append(relations[i], Connection{
				distance:   distance,
				pointIndex: j,
			})
			relations[j] = append(relations[j], Connection{
				distance:   distance,
				pointIndex: i,
			})
		}
	}

	for _, c := range relations {

		slices.SortFunc(c, func(a, b Connection) int {
			if a.distance < b.distance {
				return -1
			}
			if a.distance > b.distance {
				return 1
			}
			return 0
		})

	}

	circuits := [][]int{}
	// PAIRS := 1000
	// currentPairs := 0
	lastPair := []int{}
	for len(circuits) <= 0 || len(circuits[0]) < len(positions) {

		minDistance := math.Inf(1)
		minIndex := -1

		for i, c := range relations {
			if len(c) > 0 {
				if c[0].distance < minDistance {
					minIndex = i
					minDistance = c[0].distance
				}
			}
		}

		a := minIndex
		b := relations[a][0].pointIndex

		relations[a] = relations[a][1:]
		relations[b] = relations[b][1:]

		circuitIndexs := []int{}
		for i, c := range circuits {
			for _, x := range c {
				if x == a || x == b {
					if !slices.Contains(circuitIndexs, i) {
						circuitIndexs = append(circuitIndexs, i)
					}
				}
			}
		}
		if len(circuitIndexs) > 0 {
			secondIndex := -1
			c := circuits[circuitIndexs[0]]
			if len(circuitIndexs) > 1 {
				for _, x := range circuits[circuitIndexs[1]] {
					if !slices.Contains(c, x) {
						c = append(c, x)
					}
				}
				secondIndex = circuitIndexs[1]
			}

			if !slices.Contains(c, a) {
				c = append(c, a)
			}
			if !slices.Contains(c, b) {
				c = append(c, b)
			}

			circuits[circuitIndexs[0]] = c

			if secondIndex >= 0 {
				circuits = append(circuits[:secondIndex], circuits[secondIndex+1:]...)
			}
		} else {
			circuits = append(circuits, []int{a, b})
		}
		lastPair = []int{a, b}
		// currentPairs += 1
	}
	fmt.Println(positions[lastPair[0]].x, positions[lastPair[1]].x)
	fmt.Println(positions[lastPair[0]].x * positions[lastPair[1]].x)

	// nums := []int{}

	// for _, x := range circuits {
	// 	nums = append(nums, len(x))
	// }

	// slices.SortFunc(nums, func(a, b int) int {
	// 	return b - a
	// })

	// if len(nums) < 3 {
	// 	fmt.Println("bad")
	// } else {
	// 	fmt.Println(nums[0] * nums[1] * nums[2])
	// }

}
