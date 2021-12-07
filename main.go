package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day7"
)

func main() {
	start := time.Now()
	day7.Solve1()
	day7.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 7 took %s", elapsed)
}
