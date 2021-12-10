package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day10"
)

func main() {
	start := time.Now()
	day10.Solve1()
	day10.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 9 took %s", elapsed)
}
