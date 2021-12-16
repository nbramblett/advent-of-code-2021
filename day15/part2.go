package day15

import (
	"container/heap"
	"log"
)

func Solve2() {

	grid := ReadLines()

	log.Println(SafestPath2(grid))
}

func inBounds2(p *point, grid [][]int) bool {
	return p.i >= 0 && p.i < 5*len(grid) && p.j >= 0 && p.j < 5*len(grid[0])
}

func calcGridValue(i, j int, grid [][]int) int {
	iSize := i / len(grid)
	jSize := j / len(grid[0])
	val := grid[i%len(grid)][j%len(grid[0])]
	for i := 0; i < iSize+jSize; i++ {
		val %= 9
		val++
	}
	return val
}

// Literally just Djikstra's but bigger
func SafestPath2(grid [][]int) int {
	queue := pointHeap{}
	costs := map[point]int{}
	heap.Init(&queue)
	startingPoint := &point{0, 0, 0}
	heap.Push(&queue, startingPoint)
	previousPoint := startingPoint
	for queue.Len() != 0 {
		nextPoint := heap.Pop(&queue).(*point)
		if nextPoint.i == 5*len(grid)-1 && nextPoint.j == 5*len(grid[0])-1 {
			return nextPoint.weight
		}
		for _, adj := range adjacents(nextPoint) {
			if !inBounds2(adj, grid) || (adj.i == previousPoint.i && adj.j == previousPoint.j) {
				continue
			}
			cost, ok := costs[point{adj.i, adj.j, 0}]
			newWeight := 0
			newWeight = nextPoint.weight + calcGridValue(adj.i, adj.j, grid)
			if !ok || cost > newWeight {
				adj.weight = newWeight
				heap.Push(&queue, adj)
				costs[point{adj.i, adj.j, 0}] = newWeight
			}
		}
		previousPoint = nextPoint
	}

	return -1
}
