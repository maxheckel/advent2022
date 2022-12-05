package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type Stack struct {
	crates []rune
}

func (s *Stack) TransferMultiple(other *Stack, count int) {
	lastIndex := len(s.crates)
	appendChunk := s.crates[lastIndex-count:]
	other.crates = append(other.crates, appendChunk...)
	s.crates = s.crates[:lastIndex-count]
}

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	breakIndex := slices.Index(input, "")
	stacks := buildStacks(input)

	part1(breakIndex, input, stacks)
	// Rebuild the stacks
	stacks = buildStacks(input)
	part2(stacks, breakIndex, input)

}

func part2(stacks []*Stack, breakIndex int, input []string) {
	for i := breakIndex + 1; i < len(input); i++ {
		var count, fromIndex, toIndex int
		_, err := fmt.Sscanf(input[i], "move %d from %d to %d", &count, &fromIndex, &toIndex)
		if err != nil {
			panic(err)
		}
		stacks[fromIndex-1].TransferMultiple(stacks[toIndex-1], count)
	}

	for _, s := range stacks {
		fmt.Print(string(s.crates[len(s.crates)-1]))
	}
}

func part1(breakIndex int, input []string, stacks []*Stack) {
	for i := breakIndex + 1; i < len(input); i++ {
		var count, fromIndex, toIndex int
		_, err := fmt.Sscanf(input[i], "move %d from %d to %d", &count, &fromIndex, &toIndex)
		if err != nil {
			panic(err)
		}
		for x := 0; x < count; x++ {
			stacks[fromIndex-1].TransferMultiple(stacks[toIndex-1], 1)
		}
	}

	for _, s := range stacks {
		fmt.Print(string(s.crates[len(s.crates)-1]))
	}
	fmt.Println("")
}

func buildStacks(input []string) []*Stack {
	breakIndex := slices.Index(input, "")
	stacksInput := strings.Split(input[breakIndex-1], "  ")
	stacksCount, _ := strconv.Atoi(strings.Trim(stacksInput[len(stacksInput)-1], " "))

	stacks := make([]*Stack, stacksCount)
	// We do this in reverse order so we can pop off the end of the arrays because that's easier.
	for i := breakIndex - 2; i >= 0; i-- {
		line := input[i]

		for j := range stacks {
			crateIndex := (j * 4) + 1
			if crateIndex > len(line)-1 {
				break
			}
			if line[crateIndex] == ' ' {
				continue
			}
			if stacks[j] == nil {
				stacks[j] = &Stack{[]rune{}}
			}
			stacks[j].crates = append(stacks[j].crates, rune(line[crateIndex]))
		}
	}
	return stacks
}
