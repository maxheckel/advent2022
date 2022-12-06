package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
)

func main() {
	input, err := utils.ReadWholeFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := FindRepeatingCharsOfLength(input, 4)
	fmt.Println(part1)

	part2 := FindRepeatingCharsOfLength(input, 14)
	fmt.Println(part2)
}

func FindRepeatingCharsOfLength(input string, length int) int {
	foundChars := map[rune]*int{}
	count := 0
	for i := 0; i < len(input); i++ {
		r := rune(input[i])

		if foundChars[r] != nil {
			i = *foundChars[r]
			foundChars = map[rune]*int{}
			count = 0
			continue
		}

		// Need to reassign the location in memory otherwise i will always have the same value
		i := i
		foundChars[r] = &i
		count++

		if count == length {
			return i + 1
		}
	}
	return -1
}
