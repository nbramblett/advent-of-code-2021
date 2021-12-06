package day6

import (
	"log"
	"time"
)

// Solves Day1 as well, but is retooled to be VASTLY more space-efficient.
func Solve2() {
	inits := ReadInput()
	start := time.Now()
	cycle := make(Cycle, 9)
	for i := range cycle {
		cycle[i] = &Day{0}
	}
	for _, startingTime := range inits {
		cycle[startingTime].total++
	}
	for i := 0; i < 256; i++ {
		cycle.DoDay()
	}
	end := time.Since(start)
	log.Println(cycle.Population())
	log.Printf("Day 6 part 2 actual calculations took %s", end)
}

type Cycle []*Day

func (c Cycle) Population() (sum int64) {
	for _, day := range c {
		sum += day.total
	}
	return
}

func (c Cycle) DoDay() {
	day0Total := c[0].total
	for i := 0; i < len(c)-1; i++ {
		c[i].total = c[i+1].total
	}
	c[len(c)-1].total = day0Total
	c[6].total += day0Total
}

type Day struct {
	total int64
}
