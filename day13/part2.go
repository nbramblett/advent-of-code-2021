package day13

import (
	"fmt"
	"log"
)

func Solve2() {
	paper := ReadLines()
	ins := ReadInstructions()
	log.Println(ins)
	for _, in := range ins {
		paper.Fold(in)
	}
	PrintPaper(paper)
}

func PrintPaper(p Paper) {
	x, y := p.Size()

	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			if p[Point{j, i}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	log.Println(x, y)
}
