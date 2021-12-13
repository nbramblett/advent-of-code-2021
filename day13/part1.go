package day13

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/gohelp"
)

func Solve1() {
	paper := ReadLines()
	ins := ReadInstructions()
	paper.Fold(ins[0])
	log.Println(len(paper))
}

type Point struct {
	x, y int
}

type Paper map[Point]bool

func (p Paper) Fold(i instruction) {
	if i.axis == "y" {
		for point := range p {
			if point.y > i.value {
				newY := 2*i.value - point.y // (i.value - (point.y - i.value))
				p[Point{point.x, newY}] = true
				delete(p, point)
			}
		}
	} else {
		for point := range p {
			if point.x > i.value {
				newX := 2*i.value - point.x // (i.value - (point.x - i.value))
				p[Point{newX, point.y}] = true
				delete(p, point)
			}
		}
	}
}

func (p Paper) Size() (int, int) {
	maxX, maxY := 0, 0
	for point, v := range p {
		if !v {
			continue
		}
		if maxX < point.x {
			maxX = point.x
		}
		if maxY < point.y {
			maxY = point.y
		}
	}
	return maxX, maxY
}

func ReadLines() Paper {
	file, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	paper := Paper{}
	for scanner.Scan() {
		coords := gohelp.StringsToInts(strings.Split(scanner.Text(), ","))
		paper[Point{coords[0], coords[1]}] = true
	}
	return paper
}

type instruction struct {
	axis  string
	value int
}

func ReadInstructions() []instruction {
	file, err := os.Open("day13/instructions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []instruction{}
	for scanner.Scan() {
		ins := strings.Split(scanner.Text(), "=")
		val, err := strconv.Atoi(ins[1])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{ins[0], val})
	}
	return instructions
}
