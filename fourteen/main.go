package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"strings"
)

type Coords struct {
	x int
	y int
}

const (
	Rock = '#'
	Sand = 'o'
)

var maxHeight int

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}

func part2(input []string) {
	wall := buildWall(input)
	currentCoord := &Coords{500, 0}
	for true {
		if !moveSand(currentCoord, wall) {
			if currentCoord.x == 500 && currentCoord.y <= 0 {
				break
			}

			currentCoord.x = 500
			currentCoord.y = 0
		}
		if currentCoord.y >= maxHeight+1 {
			wall[currentCoord.y][currentCoord.x] = Sand
			currentCoord.x = 500
			currentCoord.y = 0
		}
	}
	//visualizeCorner(wall)
	total := findTotal(wall)
	fmt.Println(total)
}

func findTotal(wall [][]rune) int {
	total := 0
	for y := range wall {
		for x := range wall[y] {
			if wall[y][x] == Sand {
				total++
			}
		}
	}
	return total
}

func part1(input []string) {

	currentCoord := &Coords{500, 0}
	wall := buildWall(input)
	for true {
		if !moveSand(currentCoord, wall) {
			currentCoord.x = 500
			currentCoord.y = 0
		}
		if currentCoord.y > maxHeight {
			break
		}

	}
	//visualizeCorner(wall)
	total := findTotal(wall)
	fmt.Println(total)
}

func moveSand(currentPos *Coords, wall [][]rune) bool {

	if wall[currentPos.y+1][currentPos.x] == 0 {
		currentPos.y++
		return true
	}
	if wall[currentPos.y+1][currentPos.x-1] == 0 {
		currentPos.y++
		currentPos.x--
		return true
	}
	if wall[currentPos.y+1][currentPos.x+1] == 0 {
		currentPos.y++
		currentPos.x++
		return true
	}
	wall[currentPos.y][currentPos.x] = Sand
	return false
}

func visualizeCorner(wall [][]rune) {
	for i := range wall {
		for v := range wall[i][480:520] {
			if wall[i][480:520][v] == 0 {
				fmt.Print(" ")
			}
			fmt.Print(string(wall[i][480:520][v]))
		}
		fmt.Println("")
		if i > 10 {
			break
		}
	}
}

func buildWall(input []string) [][]rune {
	wall := make([][]rune, 700)
	for i := range wall {
		wall[i] = make([]rune, 700)
	}

	for _, line := range input {
		instructions := strings.Split(line, " -> ")
		var currentCoord *Coords
		for _, instruction := range instructions {
			var x, y int
			fmt.Sscanf(instruction, "%d,%d", &x, &y)
			if currentCoord == nil {
				currentCoord = &Coords{x, y}
				continue
			}
			xDir := 1
			yDir := 1
			if currentCoord.x > x {
				xDir = -1
			}
			if currentCoord.y > y {
				yDir = -1
			}

			for currentCoord.x != x {
				wall[currentCoord.y][currentCoord.x] = Rock
				currentCoord.x += xDir
			}
			for currentCoord.y != y {
				wall[currentCoord.y][currentCoord.x] = Rock
				currentCoord.y += yDir
			}
			wall[currentCoord.y][currentCoord.x] = Rock
			if currentCoord.y > maxHeight {
				maxHeight = currentCoord.y
			}
		}
	}
	return wall
}
