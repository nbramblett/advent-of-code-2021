package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day14"
)

func main() {
	start := time.Now()
	day14.Solve1()
	day14.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 11 took %s", elapsed)
}
