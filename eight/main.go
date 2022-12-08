package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"strconv"
)

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	trees := getTreeGrid(input)
	visibility := make([][]bool, len(trees))
	part1(visibility, trees)
	max := 0
	for y := range trees {
		for x := range trees[y] {
			left := 0
			for x1 := x - 1; x1 >= 0; x1-- {
				left++
				if trees[y][x] <= trees[y][x1] {
					break
				}
			}

			right := 0
			for x1 := x + 1; x1 <= len(trees[y])-1; x1++ {
				right++
				if trees[y][x] <= trees[y][x1] {
					break
				}
			}

			top := 0
			for y1 := y - 1; y1 >= 0; y1-- {
				top++
				if trees[y][x] <= trees[y1][x] {
					break
				}
			}

			bottom := 0
			for y1 := y + 1; y1 <= len(trees)-1; y1++ {
				bottom++
				if trees[y][x] <= trees[y1][x] {
					break
				}
			}

			score := top * bottom * left * right
			if score > max {
				max = score
			}
		}
	}
	fmt.Println(max)
}

func part1(visibility [][]bool, trees [][]int) {
	for y := range visibility {
		visibility[y] = make([]bool, len(trees[y]))
	}
	for y := range trees {
		checkRow(y, trees, visibility)
	}
	for x := range trees[0] {
		checkColumn(x, trees, visibility)
	}

	fmt.Println(visibleTrees(visibility))
}

func visibleTrees(visibility [][]bool) int {
	total := 0
	for y := range visibility {
		for x := range visibility[y] {
			if visibility[y][x] {
				total++
			}
		}
	}
	return total
}

func checkColumn(x int, trees [][]int, visibility [][]bool) {
	lastVisibility := -1
	for y := range trees {
		if lastVisibility < trees[y][x] {
			if lastVisibility < trees[y][x] {
				visibility[y][x] = true
			}
			lastVisibility = trees[y][x]
		}

	}

	lastVisibility = -1
	for y := len(trees) - 1; y >= 0; y-- {
		if lastVisibility < trees[y][x] {
			if lastVisibility < trees[y][x] {
				visibility[y][x] = true
			}
			lastVisibility = trees[y][x]
		}

	}
}

func checkRow(y int, trees [][]int, visibility [][]bool) {
	lastVisibility := -1
	for x := range trees[y] {
		if lastVisibility < trees[y][x] {
			if lastVisibility < trees[y][x] {
				visibility[y][x] = true
			}
			lastVisibility = trees[y][x]
		}

	}

	lastVisibility = -1
	for x := len(trees[y]) - 1; x >= 0; x-- {
		if lastVisibility < trees[y][x] {
			if lastVisibility < trees[y][x] {
				visibility[y][x] = true
			}
			lastVisibility = trees[y][x]
		}

	}
}

func getTreeGrid(input []string) [][]int {
	var trees [][]int

	for y, line := range input {
		trees = append(trees, []int{})
		for _, char := range line {
			intVal, _ := strconv.Atoi(string(char))
			trees[y] = append(trees[y], intVal)
		}
	}
	return trees
}
