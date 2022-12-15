package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"math"
)

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	goalRow := 10
	maxXY := 4000000
	part1(input, goalRow)
	innerEdges := map[string]bool{}
	outerEdges := map[string]bool{}
	for _, line := range input {
		var scanX, scanY, beacX, beacY int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &scanX, &scanY, &beacX, &beacY)
		distance := calcDistance(scanX, beacX, scanY, beacY)

		for y := 0; y <= distance; y++ {
			yPos := scanY + y
			xDistance := distance - y

			if scanX+xDistance > maxXY || yPos > maxXY {
				continue
			}
			outerEdges[fmt.Sprintf("%d,%d", scanX+xDistance, yPos)] = true
			outerEdges[fmt.Sprintf("%d,%d", scanX-xDistance, yPos)] = true
			innerEdges[fmt.Sprintf("%d,%d", scanX+xDistance-1, yPos)] = true
			innerEdges[fmt.Sprintf("%d,%d", scanX-xDistance+1, yPos)] = true
			yPos = scanY - y
			outerEdges[fmt.Sprintf("%d,%d", scanX+xDistance, yPos)] = true
			outerEdges[fmt.Sprintf("%d,%d", scanX-xDistance, yPos)] = true
			innerEdges[fmt.Sprintf("%d,%d", scanX+xDistance-1, yPos)] = true
			innerEdges[fmt.Sprintf("%d,%d", scanX-xDistance+1, yPos)] = true

		}
	}
	//fmt.Println(outerEdges)
	for x := 1; x < maxXY-1; x++ {
		for y := 1; y < maxXY-1; y++ {
			if (outerEdges[fmt.Sprintf("%d,%d", x+1, y+1)] || innerEdges[fmt.Sprintf("%d,%d", x+1, y+1)]) &&
				(outerEdges[fmt.Sprintf("%d,%d", x+1, y-1)] || innerEdges[fmt.Sprintf("%d,%d", x+1, y-1)]) &&
				(outerEdges[fmt.Sprintf("%d,%d", x-1, y-1)] || innerEdges[fmt.Sprintf("%d,%d", x-1, y-1)]) &&
				(outerEdges[fmt.Sprintf("%d,%d", x, y-1)] || innerEdges[fmt.Sprintf("%d,%d", x, y-1)]) &&
				(outerEdges[fmt.Sprintf("%d,%d", x, y+1)] || innerEdges[fmt.Sprintf("%d,%d", x, y+1)]) &&
				(outerEdges[fmt.Sprintf("%d,%d", x+1, y)] || innerEdges[fmt.Sprintf("%d,%d", x+1, y)]) &&
				(outerEdges[fmt.Sprintf("%d,%d", x-1, y)] || innerEdges[fmt.Sprintf("%d,%d", x-1, y)]) {
				fmt.Println(x, y)
			}
		}
	}

}

func part1(input []string, goalRow int) {
	goalRowVals := map[int]bool{}
	for _, line := range input {
		var scanX, scanY, beacX, beacY int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &scanX, &scanY, &beacX, &beacY)
		distance := calcDistance(scanX, beacX, scanY, beacY)
		distanceFromGoal := int(math.Abs(float64(scanY - goalRow)))

		if distanceFromGoal > distance {
			continue
		}

		xLength := (distance-distanceFromGoal)*2 + 1
		goalRowVals[scanX] = true
		if xLength > 1 {
			for x := 1; x <= int(math.Floor(float64(xLength/2))); x++ {
				goalRowVals[scanX-x] = true
				goalRowVals[scanX+x] = true
			}
		}
		if beacY == goalRow {
			delete(goalRowVals, beacX)
		}
	}
	fmt.Println(len(goalRowVals))
}

func calcDistance(scanX int, beacX int, scanY int, beacY int) int {
	return int(math.Abs(float64(scanX-beacX)) + math.Abs(float64(scanY-beacY)))
}
