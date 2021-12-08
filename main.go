package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day8"
)

func main() {
	start := time.Now()
	day8.Solve1()
	day8.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 7 took %s", elapsed)
}
