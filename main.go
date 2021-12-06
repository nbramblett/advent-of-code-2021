package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day4"
)

func main() {
	start := time.Now()
	day4.Solve1()
	day4.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 4 took %s", elapsed)
}
