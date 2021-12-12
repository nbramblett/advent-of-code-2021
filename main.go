package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day12"
)

func main() {
	start := time.Now()
	day12.Solve1()
	day12.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 11 took %s", elapsed)
}
