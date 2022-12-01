package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"sort"
	"strconv"
)

type Elf struct {
	Calories      []int
	TotalCalories int
}

func (e *Elf) sum() int {
	if e.TotalCalories != 0 {
		return e.TotalCalories
	}
	total := 0
	for _, i := range e.Calories {
		total += i
	}
	e.TotalCalories = total
	return e.TotalCalories
}

func main() {
	elfs := getElves()
	part1(elfs)
	part2(elfs)
}

func part2(elfs []*Elf) {
	sort.SliceStable(elfs, func(i, j int) bool {
		return elfs[i].sum() > elfs[j].sum()
	})
	fmt.Println(elfs[0].sum() + elfs[1].sum() + elfs[2].sum())
}

func part1(elfs []*Elf) {
	var maxElf *Elf
	for _, elf := range elfs {
		if maxElf == nil {
			maxElf = elf
			continue
		}
		if maxElf.sum() < elf.sum() {
			maxElf = elf
		}
	}
	fmt.Println(maxElf.sum())
}

func getElves() []*Elf {
	lines, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	var elfs []*Elf
	elf := &Elf{}
	for _, l := range lines {
		if l == "" {
			elfs = append(elfs, elf)
			elf = new(Elf)
		}
		cals, _ := strconv.Atoi(l)
		elf.Calories = append(elf.Calories, cals)
	}
	return elfs
}
