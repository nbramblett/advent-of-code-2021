package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day22"
)

func main() {
	start := time.Now()
	day22.Solve1()
	day22.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 22 took %s", elapsed)
}
