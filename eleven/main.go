package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

const (
	add = iota
	multiply
	square
)

type Monkey struct {
	items             []*Package
	operator          int
	operatorConstant  *big.Int
	testConstant      *big.Int
	positiveTestIndex int
	negativeTestIndex int
	positiveTest      *Monkey
	negativeTest      *Monkey
	inspectionCount   int64
}

func (m *Monkey) toss(item *Package, to *Monkey) {
	m.items = m.items[1:]
	to.items = append(to.items, item)
}

type Package struct {
	worry *big.Int
}

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	monkeys := getIntialMonkeyState(input, err)

	part1(monkeys)
	monkeys = getIntialMonkeyState(input, err)
	part2(monkeys)

}

func part2(monkeys []*Monkey) {
	modVal := new(big.Int)
	zero := big.NewInt(0)

	// All the divisors are prime!!!
	leastCommonMultiple := int64(1)
	for _, m := range monkeys {
		leastCommonMultiple = m.testConstant.Int64() * leastCommonMultiple
	}
	for x := 0; x < 10000; x++ {
		for _, m := range monkeys {
			for _, p := range m.items {
				m.inspectionCount++
				switch m.operator {
				case multiply:
					p.worry = p.worry.Mul(p.worry, m.operatorConstant)
				case add:
					p.worry = p.worry.Add(p.worry, m.operatorConstant)
				case square:
					p.worry = p.worry.Mul(p.worry, p.worry)
				}
				p.worry = p.worry.Mod(p.worry, big.NewInt(leastCommonMultiple))

				modVal = modVal.Mod(p.worry, m.testConstant)

				if modVal.Cmp(zero) == 0 {
					m.toss(p, m.positiveTest)
				} else {
					m.toss(p, m.negativeTest)
				}
			}
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionCount > monkeys[j].inspectionCount
	})
	fmt.Println(monkeys[0].inspectionCount * monkeys[1].inspectionCount)

}

func part1(monkeys []*Monkey) {
	for x := 0; x < 20; x++ {
		for _, m := range monkeys {
			for _, p := range m.items {
				m.inspectionCount++
				switch m.operator {
				case multiply:
					p.worry = p.worry.Mul(p.worry, m.operatorConstant)
				case add:
					p.worry = p.worry.Add(p.worry, m.operatorConstant)
				case square:
					p.worry = p.worry.Mul(p.worry, p.worry)
				}
				p.worry = p.worry.Div(p.worry, big.NewInt(3))

				modVal := new(big.Int)
				modVal = modVal.Mod(p.worry, m.testConstant)
				if modVal.Cmp(big.NewInt(0)) == 0 {
					m.toss(p, m.positiveTest)
				} else {
					m.toss(p, m.negativeTest)
				}
			}
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionCount > monkeys[j].inspectionCount
	})
	fmt.Println(monkeys[0].inspectionCount * monkeys[1].inspectionCount)

}

func getIntialMonkeyState(input []string, err error) []*Monkey {
	monkeys := []*Monkey{}
	for x := 1; x < len(input)-1; x += 7 {
		monkey := &Monkey{}
		startingItemsLine := input[x]
		parseStartingItems(startingItemsLine, monkey)
		operatorLine := input[x+1]
		err = parseOperator(err, operatorLine, monkey)
		testLine := input[x+2]
		err = parseTrueConstant(err, testLine, monkey)
		trueLine := input[x+3]
		_, err = fmt.Sscanf(trueLine, "    If true: throw to monkey %d", &monkey.positiveTestIndex)
		falseLine := input[x+4]
		_, err = fmt.Sscanf(falseLine, "    If false: throw to monkey %d", &monkey.negativeTestIndex)
		if err != nil {
			panic(err)
		}
		monkeys = append(monkeys, monkey)
	}
	for _, m := range monkeys {
		m.positiveTest = monkeys[m.positiveTestIndex]
		m.negativeTest = monkeys[m.negativeTestIndex]
	}

	return monkeys
}

func parseTrueConstant(err error, testLine string, monkey *Monkey) error {
	var testConstant int64
	_, err = fmt.Sscanf(testLine, "  Test: divisible by %d", &testConstant)
	monkey.testConstant = big.NewInt(testConstant)
	return err
}

func parseStartingItems(startingItemsLine string, monkey *Monkey) {
	startingItemsLine = strings.TrimLeft(startingItemsLine, "  Starting items: ")
	startingItemsStrings := strings.Split(startingItemsLine, ", ")
	for _, item := range startingItemsStrings {
		val, _ := strconv.Atoi(item)
		monkey.items = append(monkey.items, &Package{worry: big.NewInt(int64(val))})
	}
}

func parseOperator(err error, operatorLine string, monkey *Monkey) error {
	var operator rune
	var operatorConstant int64
	_, err = fmt.Sscanf(operatorLine, "  Operation: new = old %c %d", &operator, &operatorConstant)
	monkey.operatorConstant = big.NewInt(operatorConstant)
	if err != nil {
		_, err = fmt.Sscanf(operatorLine, "  Operation: new = old %c old", &operator)
		monkey.operator = square
	} else {
		switch operator {
		case '*':
			monkey.operator = multiply
		case '+':
			monkey.operator = add
		}
	}
	return err
}
