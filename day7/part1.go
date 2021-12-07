package day7

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/gohelp"
)

var distanceMetric func(x, y int) int

func Solve1() {
	vals := ReadInput()
	distanceMetric = func(x, y int) int {
		return int(math.Abs(float64(x - y)))
	}
	od := optimizeDistance(vals)
	log.Println(od)
}

func optimizeDistance(vals []int) int {
	min, max := gohelp.MinMax(vals...)
	optimalDistance := math.MaxInt32
	for i := min; i <= max; i++ {
		od := totalDistance(i, vals)
		if optimalDistance > od {
			optimalDistance = od
		}
	}
	return optimalDistance
}

func totalDistance(target int, vals []int) int {
	sum := 0
	for _, val := range vals {
		sum += distanceMetric(target, val)
	}
	return sum
}

func ReadInput() []int {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	for scanner.Scan() {
		line := scanner.Text()
		return gohelp.StringsToInts(strings.Split(line, ","))
	}
	panic("no line!")
}
