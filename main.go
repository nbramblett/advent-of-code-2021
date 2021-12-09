package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day9"
)

func main() {
	start := time.Now()
	day9.Solve1()
	day9.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 7 took %s", elapsed)
}
