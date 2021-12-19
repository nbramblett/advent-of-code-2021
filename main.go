package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day19"
)

func main() {
	start := time.Now()
	day19.Solve1()
	//day19.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 18 took %s", elapsed)
}
