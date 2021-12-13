package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day13"
)

func main() {
	start := time.Now()
	day13.Solve1()
	day13.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 11 took %s", elapsed)
}
