package day7

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/nbramblett/advent-of-code-2021/util"
)

var distanceMetric func(x, y int) int

func Solve1() {
	now := time.Now()
	vals := ReadInput()
	distanceMetric = func(x, y int) int {
		return int(math.Abs(float64(x - y)))
	}
	od := optimizeDistance(vals)
	elapsed := time.Since(now)
	log.Println(od)
	log.Printf("Day 6 part 1 actual calculations took %s", elapsed)
}

func optimizeDistance(vals []int) int {
	min, max := util.MinMax(vals...)
	optimalDistance := math.MaxInt32
	optimalTarget := 0
	for i := min; i <= max; i++ {
		od := totalDistance(i, vals)
		if optimalDistance > od {
			optimalDistance = od
			optimalTarget = i
		}
	}
	log.Println(optimalTarget)
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
	file, err := os.Open("day7/input2.txt")
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
