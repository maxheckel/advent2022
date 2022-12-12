package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
)

type Node struct {
	Children map[*Node]*Node
	Val      rune
}

var StartNode *Node
var EndNode *Node

func main() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	buildNodeList(input)
	queue := []*Node{EndNode}
	visited := map[*Node]int{}
	var shortest *Node
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if shortest == nil && current.Val == 'a' {
			shortest = current
		}
		for _, child := range current.Children {
			if visited[child] > 0 {
				continue
			}

			visited[child] = visited[current] + 1
			queue = append(queue, child)
		}
	}
	fmt.Println(visited[StartNode])
	fmt.Println(visited[shortest])
}

func buildNodeList(input []string) []*Node {
	terrain := make([][]*Node, len(input))
	for y, line := range input {
		for _, r := range line {
			newNode := &Node{
				Children: map[*Node]*Node{},
			}
			if r == 'S' {
				r = 'a'
				StartNode = newNode
			}
			if r == 'E' {
				r = 'z'
				EndNode = newNode
			}
			newNode.Val = r
			terrain[y] = append(terrain[y], newNode)
		}
	}

	for y := range terrain {
		for x := range terrain[y] {
			nodeList := []*Node{}
			if x > 0 {
				nodeList = append(nodeList, terrain[y][x-1])
			}
			if x != len(terrain[y])-1 {
				nodeList = append(nodeList, terrain[y][x+1])
			}
			if y > 0 {
				nodeList = append(nodeList, terrain[y-1][x])
			}
			if y != len(terrain)-1 {
				nodeList = append(nodeList, terrain[y+1][x])
			}
			for i, n := range nodeList {
				if terrain[y][x].Val <= n.Val+1 {
					terrain[y][x].Children[nodeList[i]] = nodeList[i]
				}
			}
		}
	}
	var res []*Node
	for y := range terrain {
		res = append(res, terrain[y]...)
	}
	return res
}
