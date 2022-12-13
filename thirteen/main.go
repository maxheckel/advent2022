package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"sort"
	"strconv"
)

type Res int

const (
	False Res = iota
	True
	NoDecision
)

func compare(left, right []interface{}) Res {
	if len(left) == 0 && len(right) > 0 {
		return True
	}
	if len(right) == 0 && len(left) > 0 {
		return False
	}
	for i := range left {

		// right runs out first
		if i > len(right)-1 {
			return False
		}
		// both are arrays
		if leftVal, ok := left[i].([]interface{}); ok {
			if rightVal, ok := right[i].([]interface{}); ok {
				cmpResult := compare(leftVal, rightVal)
				if cmpResult != NoDecision {
					return cmpResult
				}
			}
		}

		// leftVal is array, rightVal is int
		if leftVal, ok := left[i].([]interface{}); ok {
			if rightVal, ok := right[i].(int); ok {
				res := compare(leftVal, []interface{}{rightVal})
				if res != NoDecision {
					return res
				}
			}
		}

		// rightVal is array, leftVal is int
		if leftVal, ok := left[i].(int); ok {
			if rightVal, ok := right[i].([]interface{}); ok {
				res := compare([]interface{}{leftVal}, rightVal)
				if res != NoDecision {
					return res
				}
			}
		}

		// Both are ints
		if leftVal, ok := left[i].(int); ok {
			if rightVal, ok := right[i].(int); ok {
				if rightVal < leftVal {
					return False
				}
				if leftVal < rightVal {
					return True
				}
			}
		}

		// left runs out first
		if i == len(left)-1 && i < len(right)-1 {
			return True
		}

	}

	return NoDecision
}

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}

func part2(input []string) {
	var packets [][]interface{}
	divider1 := []interface{}{
		6,
	}
	divider2 := []interface{}{
		2,
	}
	packets = append(packets, divider1)
	packets = append(packets, divider2)
	for _, line := range input {
		if line == "" {
			continue
		}
		a, _ := parseLine(line, 0)
		// My input parser incorrectly adds an array at the top level
		if nest, ok := a[0].([]interface{}); ok {
			a = nest
		}
		packets = append(packets, a)
	}
	sort.SliceStable(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == True
	})

	// comparing interfaces is too hard, copy the output and look for the answer lines then multiply them manually
	for _, packet := range packets {
		fmt.Println(packet)
	}
}

func part1(input []string) {
	var a []interface{}
	var b []interface{}
	total := 0
	setIndex := 1
	for _, line := range input {
		if line == "" {
			res := compare(a, b)
			fmt.Println(fmt.Sprintf("Comparision %d was %t", setIndex, res))
			if res == True {
				total += setIndex
			}
			setIndex++
			a = nil
			b = nil
			continue
		}
		if a == nil {
			a, _ = parseLine(line, 0)
			// My input parser incorrectly adds an array at the top level
			if nest, ok := a[0].([]interface{}); ok {
				a = nest
			}
			continue
		}
		if b == nil {
			b, _ = parseLine(line, 0)
			// My input parser incorrectly adds an array at the top level
			if nest, ok := b[0].([]interface{}); ok {
				b = nest
			}
		}

	}
	fmt.Println(total)
}

func parseLine(line string, index int) ([]interface{}, int) {

	resp := []interface{}{}
	for ; index < len(line); index++ {
		c := line[index]
		if c == '[' {
			var innerArray []interface{}
			innerArray, index = parseLine(line, index+1)
			resp = append(resp, innerArray)
			continue
		}
		if c == ']' {
			return resp, index
		}
		if c == ',' {
			continue
		}
		numParse := index
		str := ""
		for line[numParse] >= 48 && line[numParse] <= 57 {
			str = fmt.Sprintf("%s%c", str, line[numParse])
			numParse++
		}
		intVal, _ := strconv.Atoi(str)
		resp = append(resp, intVal)
	}
	return resp, index
}
