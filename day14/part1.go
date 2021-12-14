package day14

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

const line = "KKOSPHCNOCHHHSPOBKVF"

func Solve1() {
	steps := ReadInput()
	l := line
	m := countPairs(l)
	counts := map[string]int{}
	for _, r := range line {
		counts[string(r)]++
	}
	for i := 0; i < 40; i++ {
		maps := []map[string]int{}
		for _, step := range steps {
			newMap := applyToChain(l, step[0], step[1], m, counts)
			maps = append(maps, newMap)
			delete(m, step[0])
		}
		for _, ms := range maps {
			for k, v := range ms {
				m[k] += v
			}
		}
	}
	min, max := math.MaxInt64, 0
	for _, c := range counts {
		if min > c {
			min = c
		}
		if max < c {
			max = c
		}
	}
	log.Println(max, min, max-min)
}

func countPairs(s string) map[string]int {
	m := map[string]int{}

	for i := range s {
		if i < len(s)-1 {
			m[s[i:i+2]] += 1
		}
	}
	return m
}

func applyToChain(l string, start string, end string, pairs, counts map[string]int) map[string]int {
	if pairs[start] == 0 {
		return nil
	}
	nm := map[string]int{}
	st := pairs[start]
	nm[start[0:1]+end] = st
	nm[end+start[1:]] = st
	counts[end] += st
	return nm
}

func ReadInput() [][]string {
	file, err := os.Open("day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := [][]string{}
	for scanner.Scan() {
		instructions = append(instructions, strings.Split(scanner.Text(), " -> "))
	}
	return instructions
}
