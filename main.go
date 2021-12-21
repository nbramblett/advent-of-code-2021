package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day21"
)

func main() {
	start := time.Now()
	day21.Solve1()
	day21.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 20 took %s", elapsed)
}
