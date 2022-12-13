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
	var a []interface{}
	var b []interface{}
	for _, line := range input {
		if line == "" {
			a = nil
			b = nil
			fmt.Println("compare")
			continue
		}
		if a == nil {
			a, _ = parseLine(line, 0)
			// My input parser incorrectly adds an array at the top level
			if nest, ok := a[0].([]interface{}); ok {
				a = nest
			}
			fmt.Println(a)
			continue
		}
		if b == nil {
			b, _ = parseLine(line, 0)
			// My input parser incorrectly adds an array at the top level
			if nest, ok := b[0].([]interface{}); ok {
				b = nest
			}
			fmt.Println(b)
		}
	}
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
