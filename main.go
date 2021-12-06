package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day5"
)

func main() {
	start := time.Now()
	day5.Solve1()
	day5.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 5 took %s", elapsed)
}
