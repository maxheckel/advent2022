package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
)

func main() {
	lines, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	part1(lines)
	part2(lines)
}

func part2(lines []string) {
	scores := []rune{}
TOP:
	for i := 0; i < len(lines)/3; i++ {
		elf1 := lines[i*3]
		elf2 := lines[i*3+1]
		elf3 := lines[i*3+2]
		sharedMap := map[rune]bool{}
		elf1Map := map[rune]bool{}
		for _, e1 := range elf1 {
			elf1Map[e1] = true
		}
		for _, e2 := range elf2 {
			sharedMap[e2] = elf1Map[e2]
		}
		for _, e3 := range elf3 {
			if sharedMap[e3] {
				scores = append(scores, e3)
				continue TOP
			}
		}
	}
	fmt.Println(runesToScores(scores))
}

func part1(lines []string) {
	scores := []rune{}
TOP:
	for _, line := range lines {
		half1 := line[:len(line)/2]
		half2 := line[len(line)/2:]
		left := map[rune]bool{}
		for _, r := range half1 {
			left[r] = true
		}
		for _, r := range half2 {
			if left[r] {
				scores = append(scores, r)
				continue TOP
			}
		}
	}

	total := runesToScores(scores)

	fmt.Println(total)
}

func runesToScores(scores []rune) int {
	total := 0
	for _, s := range scores {
		score := int(s) - 96
		if score < 0 {
			score = int(s) - 64 + 26
		}
		total += score
	}
	return total
}
