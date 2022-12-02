package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"strings"
)

var opponentInputs = []rune{'A', 'B', 'C'}
var myInputs = []rune{'X', 'Y', 'Z'}

const (
	lose = iota
	draw
	win
)

func main() {
	inputs, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	part1(inputs)
	part2(inputs)

}

func part1(inputs []string) {
	score := 0
	for _, gameStr := range inputs {
		game := strings.Split(gameStr, " ")
		opponentPlay := indexOfRune(rune(game[0][0]), opponentInputs)
		myPlay := indexOfRune(rune(game[1][0]), myInputs)
		score += myPlay + 1
		switch whoWon(myPlay, opponentPlay) {
		case win:
			score += 6
		case draw:
			score += 3
		}
	}

	fmt.Println(score)
}

func part2(inputs []string) {
	score := 0
	for _, gameStr := range inputs {
		game := strings.Split(gameStr, " ")
		opponentPlay := indexOfRune(rune(game[0][0]), opponentInputs)
		var myPlay int
		switch indexOfRune(rune(game[1][0]), myInputs) {
		case win:
			score += 6
			myPlay = opponentPlay + 1
			if myPlay == 3 {
				myPlay = 0
			}
		case lose:
			myPlay = opponentPlay - 1
			if myPlay == -1 {
				myPlay = 2
			}
		case draw:
			score += 3
			myPlay = opponentPlay
		}
		score += myPlay + 1

	}

	fmt.Println(score)
}

func whoWon(myPlay, opponentPlay int) int {
	// Draw
	if myPlay == opponentPlay {
		return draw
	}

	// I win
	if opponentPlay+1 == myPlay || (opponentPlay == 2 && myPlay == 0) {
		return win
	}

	return lose

}

func indexOfRune(r rune, list []rune) int {
	for i, check := range list {
		if check == r {
			return i
		}
	}
	return -1
}
