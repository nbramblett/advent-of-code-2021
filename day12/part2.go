package day12

import (
	"fmt"
	"log"
	"strings"
)

// Solves the duplication problem
var pathMap = map[string]bool{}

func Solve2() {
	cave := ReadInput()
	for l := range cave {
		if strings.ToUpper(l) != l && l != "start" && l != "end" {
			Path2(cave, cave["start"], map[string]bool{}, l, "")
		}
	}

	log.Println(len(pathMap))
}

func Path2(cave NodeSet, currentNode *Node, visited map[string]bool, doubled string, path string) {
	if currentNode.label == "end" {
		path := fmt.Sprintf("%s--end", path)
		pathMap[path] = true
		return
	}
	if strings.ToUpper(currentNode.label) != currentNode.label {
		if currentNode.label == doubled {
			doubled = ""
		} else {
			visited[currentNode.label] = true
		}
	}
	for _, neighbor := range currentNode.linked {
		if !visited[neighbor.label] {
			nv := copyVisited(visited)
			Path2(cave, neighbor, nv, doubled, fmt.Sprintf("%s-%s", path, currentNode.label))
		}
	}
}
