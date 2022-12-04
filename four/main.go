package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"strconv"
	"strings"
)

type Section struct {
	min int
	max int
}

func (s Section) Contains(otherSection *Section) bool {
	return s.min <= otherSection.min && s.max >= otherSection.max
}
func (s Section) Overlaps(otherSection *Section) bool {
	return (otherSection.min <= s.min && otherSection.max >= s.min) || (otherSection.max >= s.max && otherSection.min <= s.max)
}

func NewFromRow(row string) *Section {
	split := strings.Split(row, "-")
	section := &Section{}
	section.min, _ = strconv.Atoi(split[0])
	section.max, _ = strconv.Atoi(split[1])
	return section
}

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}

func part1(input []string) {
	count := 0
	for _, line := range input {
		pairs := strings.Split(line, ",")
		s1 := NewFromRow(pairs[0])
		s2 := NewFromRow(pairs[1])
		if s1.Contains(s2) || s2.Contains(s1) {
			count++
		}
	}
	fmt.Println(count)
}

func part2(input []string) {
	count := 0
	for _, line := range input {
		pairs := strings.Split(line, ",")
		s1 := NewFromRow(pairs[0])
		s2 := NewFromRow(pairs[1])
		if s1.Overlaps(s2) || s2.Overlaps(s1) {
			count++
		}
	}
	fmt.Println(count)
}
