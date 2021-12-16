package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/util"
)

func Solve1() {
	inits := ReadInput()
	colony := make(Colony, 0, 10000)
	for _, startingTime := range inits {
		colony = append(colony, &Fish{startingTime})
	}
	colony.Days(79)
	fmt.Println(len(colony))
}

const cycle = 7

type Colony []*Fish

func (c *Colony) Days(days int) {
	for i := 0; i < days; i++ {
		for _, fish := range *c {
			newFish := fish.Day()
			if newFish != nil {
				*c = append(*c, newFish)
			}
		}
	}
}

type Fish struct {
	daysLeft int
}

func (fish *Fish) Day() *Fish {
	fish.daysLeft--
	if fish.daysLeft == 0 {
		fish.daysLeft = cycle
		return &Fish{cycle + 2}
	}
	return nil
}

func ReadInput() []int {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	for scanner.Scan() {
		line := scanner.Text()
		return util.StringsToInts(strings.Split(line, ","))
	}
	panic("no line!")
}
