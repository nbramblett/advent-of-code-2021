package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day6"
)

func main() {
	start := time.Now()
	day6.Solve1()
	day6.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 6 took %s", elapsed)
}
