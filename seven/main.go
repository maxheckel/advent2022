package main

import (
	"fmt"
	"github.com/maxheckel/advent2022/utils"
	"math"
)

const (
	File = iota
	Directory
)

type Node struct {
	parent   *Node
	children []*Node
	nodeType int
	name     string
	size     int32
	pwd      string
}

var rootNode *Node
var currentNode *Node
var totalSize int32

func main() {
	buildNodes()

	part1()

	part2()

}

func part1() {
	getSizesLessThan10k(rootNode)
	fmt.Println(totalSize)
}

func part2() {
	flatNodes := []*Node{}
	flatNodes = addFlatNodes(rootNode)
	var minNode *Node
	minDiff := float64(10000000000000)
	goalSize := float64(30000000) - (float64(70000000) - float64(rootNode.size))
	for _, n := range flatNodes {
		if float64(n.size) < goalSize {
			continue
		}
		thisDiff := math.Abs(float64(n.size) - goalSize)
		if thisDiff < minDiff {
			minDiff = thisDiff
			minNode = n
		}
	}
	fmt.Println(minNode.size)
}

func buildNodes() {
	input, err := utils.ReadInputLines("input.txt")
	if err != nil {
		panic(err)
	}
	rootNode = &Node{
		parent:   nil,
		children: []*Node{},
		nodeType: Directory,
		name:     "",
		size:     0,
		pwd:      "/",
	}
	for i := 0; i < len(input); i++ {
		output := input[i]
		switch output[0] {
		case '$':
			var cmd string
			var args string
			fmt.Sscanf(output, "$ %s %s", &cmd, &args)
			if cmd == "cd" {
				handleCD(args)
				continue
			}
		default:
			handleNewNode(output)
		}
	}
}

func addFlatNodes(current *Node) []*Node {
	var fullList []*Node
	if current.nodeType == Directory {
		fullList = append(fullList, current)
	}

	for _, child := range current.children {
		if child.nodeType == Directory {
			fullList = append(fullList, addFlatNodes(child)...)
		}

	}
	return fullList
}

func (n *Node) addNode(newNode *Node) {
	n.children = append(n.children, newNode)
	if newNode.nodeType == Directory {
		return
	}
	checkNode := n
	for checkNode != nil {
		checkNode.size += newNode.size
		checkNode = checkNode.parent
	}
}

func getSizesLessThan10k(node *Node) {
	if node.nodeType == Directory && node.size <= 100000 {
		totalSize += node.size
	}
	if node.children != nil && len(node.children) > 0 {
		for _, child := range node.children {
			getSizesLessThan10k(child)
		}
	}
}

func handleNewNode(output string) {
	var fileSize int32
	var nodeName string
	fmt.Sscanf(output, "%d %s", &fileSize, &nodeName)
	nodeType := File
	// This means it's a child file of the currentNode
	if fileSize == 0 {
		nodeType = Directory
		fmt.Sscanf(output, "dir %s", &nodeName)
	}

	// Check so we don't add the same file multiple times
	for _, child := range currentNode.children {
		if child.name == nodeName && child.nodeType == nodeType {
			return
		}
	}
	currentNode.addNode(&Node{
		parent:   currentNode,
		children: []*Node{},
		nodeType: nodeType,
		name:     nodeName,
		size:     fileSize,
		pwd:      currentNode.pwd + "/" + nodeName,
	})
}

func handleCD(args string) {

	switch args {
	case "/":
		currentNode = rootNode
	case "..":
		currentNode = currentNode.parent
	default:
		for _, child := range currentNode.children {
			if child.name == args && child.nodeType == Directory {
				currentNode = child
			}
		}
	}
}
