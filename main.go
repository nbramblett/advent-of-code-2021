package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day17"
)

func main() {
	start := time.Now()
	day17.Solve1()
	//day16.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 11 took %s", elapsed)
}
