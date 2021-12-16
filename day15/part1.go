package day15

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/gohelp"
)

func Solve1() {

	grid := ReadLines()

	log.Println(SafestPath(grid))
}

type point struct {
	i, j   int
	weight int
}

type pointHeap []*point

// Implement heap interface
func (p pointHeap) Less(i, j int) bool {
	return p[i].weight < p[j].weight
}

func (p pointHeap) Swap(i, j int) {
	p2 := p[i]
	p[i] = p[j]
	p[j] = p2
}

func (p *pointHeap) Push(x interface{}) {
	p2 := x.(*point)
	*p = append(*p, p2)
}

func (p *pointHeap) Pop() interface{} {
	p2 := (*p)[len(*p)-1]
	*p = (*p)[0 : len(*p)-1]
	return p2
}

func (p pointHeap) Len() int {
	return len(p)
}

func adjacents(p *point) []*point {
	lp := []*point{}
	lp = append(lp, &point{p.i - 1, p.j, 0}, &point{p.i + 1, p.j, 0}, &point{p.i, p.j - 1, 0}, &point{p.i, p.j + 1, 0})
	return lp
}

func inBounds(p *point, grid [][]int) bool {
	return p.i >= 0 && p.i < len(grid) && p.j >= 0 && p.j < len(grid[p.i])
}

// Literally just Djikstra's
func SafestPath(grid [][]int) int {
	queue := pointHeap{}
	costs := map[point]int{}
	heap.Init(&queue)
	startingPoint := &point{0, 0, 0}
	heap.Push(&queue, startingPoint)
	previousPoint := startingPoint
	for queue.Len() != 0 {
		nextPoint := heap.Pop(&queue).(*point)
		if nextPoint.i == len(grid)-1 && nextPoint.j == len(grid[0])-1 {
			return nextPoint.weight
		}
		for _, adj := range adjacents(nextPoint) {
			if !inBounds(adj, grid) || (adj.i == previousPoint.i && adj.j == previousPoint.j) {
				continue
			}
			cost, ok := costs[point{adj.i, adj.j, 0}]
			newWeight := nextPoint.weight + grid[adj.i][adj.j]
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

func ReadLines() [][]int {
	file, err := os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	grid := [][]int{}
	for scanner.Scan() {
		points := gohelp.StringsToInts(strings.Split(scanner.Text(), ""))
		grid = append(grid, points)
	}
	return grid
}
