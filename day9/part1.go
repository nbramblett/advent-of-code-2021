package day9

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Solve1() {

}

func ReadInput() [][]string {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	lines := [][]string{}
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), " | "))
	}
	return lines
	panic("no line!")
}
