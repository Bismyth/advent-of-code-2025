package main


import (
	"fmt"
	"strings"
	"os"
)

func part1(lines []string) {
	beams := []int{}
	splits := 0
	for index, l := range lines {
		if index == 0 {
			for i, c := range l {
				if c == 'S' {
					beams = append(beams, i)
				}
			}
		} else {
			lineLength := len(l)
			newBeams := []int{}
			currentBeamIndex := 0
			
			for i, c := range l {
				if i > beams[currentBeamIndex] {
					if (len(beams) - 1) <= currentBeamIndex {
						break;
					}
					currentBeamIndex += 1
				}
				if i == beams[currentBeamIndex] {
					if c == '.' {
						if len(newBeams) <= 0 {
							newBeams = append(newBeams, i)
						} else if newBeams[len(newBeams) - 1] != i {
							newBeams = append(newBeams, i)
						} 
						
					} else if c == '^' {
						splits += 1
						if i > 0 {
							if len(newBeams) <= 0 {
								newBeams = append(newBeams, i-1)
							} else if newBeams[len(newBeams) - 1] != i-1 {
								newBeams = append(newBeams, i-1)
							}
						}
						if i < lineLength-1 {
							newBeams = append(newBeams, i+1)
						}
					}
				}
			}
			//fmt.Println(newBeams)
			beams = newBeams
		}
		
		
		
	}
	fmt.Println(splits)
}

func part2(lines []string) {
	timelines := []int{}
	beams := []int{}
	for index, l := range lines {
		if index == 0 {
		
			
			timelines = make([]int, len(l))
			for i, c := range l {
				if c == 'S' {
					timelines[i] = 1
					beams = append(beams, i)
				} else {
					timelines[i] = 0
				}
			}
		} else {
			lineLength := len(l)
			newTimelines := make([]int, len(l))
			currentBeamIndex := 0
			newBeams := []int{}
			for i, c := range l {
				if i > beams[currentBeamIndex] {
					if (len(beams) - 1) <= currentBeamIndex {
						break;
					}
					currentBeamIndex += 1
				}
				if i == beams[currentBeamIndex] {
					if c == '.' {
						if len(newBeams) <= 0 {
							newBeams = append(newBeams, i)
						} else if newBeams[len(newBeams) - 1] != i {
							newBeams = append(newBeams, i)
						} 
						newTimelines[i] += timelines[i]
						
					} else if c == '^' {
						if i > 0 {
							if len(newBeams) <= 0 {
								newBeams = append(newBeams, i-1)
							} else if newBeams[len(newBeams) - 1] != i-1 {
								newBeams = append(newBeams, i-1)
							}
							newTimelines[i-1] += timelines[i]
						}
						if i < lineLength-1 {
							newBeams = append(newBeams, i+1)
						 	newTimelines[i+1] += timelines[i]
						}
					}
				}
			}
			if len(newTimelines) > 0 {
				timelines = newTimelines
			}
			beams = newBeams
		}
		
		
	}
	total := 0
	for _, n := range timelines {
		total += n
	}
	fmt.Println(total)
}


func main() {

	dat, err := os.ReadFile("./day-06-input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n") 
	//part1(lines)
	part2(lines)
}
