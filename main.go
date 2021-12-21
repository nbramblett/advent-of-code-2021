package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day20"
)

func main() {
	start := time.Now()
	day20.Solve1()
	day20.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 20 took %s", elapsed)
}
