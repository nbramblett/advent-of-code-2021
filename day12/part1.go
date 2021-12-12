package day12

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Solve1() {
	cave := ReadInput()
	paths := Path(cave, cave["start"], map[string]bool{})
	log.Println(paths)
}

func Path(cave NodeSet, currentNode *Node, visited map[string]bool) int {
	if currentNode.label == "end" {
		return 1
	}
	if strings.ToUpper(currentNode.label) != currentNode.label {
		visited[currentNode.label] = true
	}
	sum := 0
	for _, neighbor := range currentNode.linked {
		if !visited[neighbor.label] {

			nv := copyVisited(visited)
			sum += Path(cave, neighbor, nv)
		}
	}
	return sum
}

type Node struct {
	label  string
	linked []*Node
}

type NodeSet map[string]*Node

func copyVisited(m map[string]bool) map[string]bool {
	newMap := map[string]bool{}
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func (n NodeSet) AddNodes(label, label2 string) {
	node, ok := n[label]
	linkedNode, lok := n[label2]
	if !lok {
		linkedNode = &Node{label: label2}
	}
	if !ok {
		node = &Node{label: label}
	}
	node.linked = append(node.linked, linkedNode)
	linkedNode.linked = append(linkedNode.linked, node)

	n[label] = node
	n[label2] = linkedNode
}

func ReadInput() NodeSet {
	file, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	nodes := NodeSet{}
	for scanner.Scan() {
		caves := strings.Split(scanner.Text(), "-")
		nodes.AddNodes(caves[0], caves[1])
	}
	return nodes
}
