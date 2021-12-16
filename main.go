package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day15"
)

func main() {
	start := time.Now()
	day15.Solve1()
	day15.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 11 took %s", elapsed)
}
