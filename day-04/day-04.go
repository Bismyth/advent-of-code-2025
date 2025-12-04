package main

import (
	"fmt"
	"os"
	"strings"
)

func getNeighbours(grid [][]bool, x int, y int) int {
	res := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if grid[x+i][y+j] {
				res += 1
			}

		}
	}

	return res

}

func main() {

	grid := [][]bool{}

	dat, err := os.ReadFile("./day-04-input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	gridSize := len(lines)
	grid = append(grid, make([]bool, gridSize+2))

	for _, l := range lines {
		row := make([]bool, gridSize+2)
		for i, c := range l {
			row[i+1] = c == '@'
		}
		grid = append(grid, row)
	}

	grid = append(grid, make([]bool, gridSize+2))

	// for _, l := range grid {
	// 	for _, b := range l {
	// 		if b {
	// 			fmt.Print("@")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	answer := 0

	didRemove := true

	for didRemove {
		didRemove = false
		for i := 1; i < gridSize+1; i++ {
			for j := 1; j < gridSize+1; j++ {
				if grid[i][j] && getNeighbours(grid, i, j) < 4 {
					answer += 1
					grid[i][j] = false
					didRemove = true
				}
			}
		}
	}

	fmt.Println(answer)
}
