package day9

import (
	"log"
	"math"
)

func Solve2() {
	grid := ReadInput()
	lowPoints := LowCoords(grid)
	log.Println(Basins(grid, lowPoints))
}

func LowCoords(grid [][]int) [][]int {
	lowVals := [][]int{}
	for i := range grid {
		for k, v := range grid[i] {
			lowest := true
			if k > 0 && grid[i][k-1] <= v {
				lowest = false
			}
			if k < len(grid[i])-1 && grid[i][k+1] <= v {
				lowest = false
			}
			if i > 0 && grid[i-1][k] <= v {
				lowest = false
			}
			if i < len(grid)-1 && grid[i+1][k] <= v {
				lowest = false
			}
			if lowest {
				lowVals = append(lowVals, []int{i, k})
			}
		}
	}
	return lowVals
}

// Assess the size of each individual basin and keep a running tally of the top 3
func Basins(grid [][]int, lowPoints [][]int) []int {
	top3 := make([]int, 3)
	for _, coords := range lowPoints {
		size := Basin(grid, coords[0], coords[1])
		smallestI := 0
		smallestV := math.MaxInt32
		for i, v := range top3 {
			if v < smallestV {
				smallestV = v
				smallestI = i
			}
		}
		if size > smallestV {
			top3[smallestI] = size
		}
	}
	return top3
}

type Vector struct {
	i, k int
}

// Literally just breadth-first search. Start at a low point and go outward, stopping when you hit 9 or an OOB (and ignoring repeats ofc)
func Basin(grid [][]int, lowI, lowK int) int {
	queue := []Vector{{lowI, lowK}}
	visited := map[Vector]bool{}
	for len(queue) != 0 {
		i, k := queue[0].i, queue[0].k
		visited[Vector{i, k}] = true
		queue = queue[1:]
		if k > 0 && grid[i][k-1] != 9 && !visited[Vector{i, k - 1}] {
			queue = append(queue, Vector{i, k - 1})
		}
		if k < len(grid[i])-1 && grid[i][k+1] != 9 && !visited[Vector{i, k + 1}] {
			queue = append(queue, Vector{i, k + 1})
		}
		if i > 0 && grid[i-1][k] != 9 && !visited[Vector{i - 1, k}] {
			queue = append(queue, Vector{i - 1, k})
		}
		if i < len(grid)-1 && grid[i+1][k] != 9 && !visited[Vector{i + 1, k}] {
			queue = append(queue, Vector{i + 1, k})
		}
	}
	return len(visited)
}
