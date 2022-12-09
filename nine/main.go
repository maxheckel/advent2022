package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"math"
)

type Coords struct {
	X int
	Y int
}

func (tail *Coords) MoveTo(head *Coords) {
	// If the head is right next to or directly on top of the tail then we do nothing
	diffInX := math.Abs(float64(tail.X - head.X))
	diffInY := math.Abs(float64(tail.Y - head.Y))
	if diffInY <= 1 && diffInX <= 1 {
		return
	}
	//fmt.Println(fmt.Sprintf("Tail moving (%s) to meet head (%s)", tail.String(), head.String()))
	xDirection := 1
	yDirection := 1
	if head.Y < tail.Y {
		yDirection = -1
	}
	if head.X < tail.X {
		xDirection = -1
	}
	if head.X == tail.X {
		tail.Y += 1 * yDirection
		return
	}
	if head.Y == tail.Y {
		tail.X += 1 * xDirection
		return
	}

	// Diagonal
	tail.X += 1 * xDirection
	tail.Y += 1 * yDirection
}

func (c *Coords) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	part1(input)
	coords := make([]Coords, 10)
	tailPositions := map[string]bool{}
	tailPositions[coords[9].String()] = true
	for _, cmd := range input {
		var direction string
		var amount int
		fmt.Sscanf(cmd, "%s %d", &direction, &amount)
		for x := 0; x < amount; x++ {

			switch direction {
			case "R":
				coords[0].X++
			case "U":
				coords[0].Y++
			case "D":
				coords[0].Y--
			case "L":
				coords[0].X--
			}
			for i := 1; i < len(coords); i++ {
				coords[i].MoveTo(&coords[i-1])
			}
			tailPositions[coords[9].String()] = true
		}
	}
	fmt.Println(len(tailPositions))
}

func part1(input []string) {
	// Memoize where the tail has been with x,y being the keys
	tailPositions := map[string]bool{}
	headPosition := &Coords{0, 0}
	tailPosition := &Coords{0, 0}
	tailPositions[tailPosition.String()] = true
	for _, cmd := range input {
		var direction string
		var amount int
		fmt.Sscanf(cmd, "%s %d", &direction, &amount)
		for x := 0; x < amount; x++ {

			switch direction {
			case "R":
				headPosition.X++
			case "U":
				headPosition.Y++
			case "D":
				headPosition.Y--
			case "L":
				headPosition.X--
			}
			tailPosition.MoveTo(headPosition)
			tailPositions[tailPosition.String()] = true
		}
	}

	fmt.Println(len(tailPositions))
}
