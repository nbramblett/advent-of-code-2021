package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	nums, boards := ReadInput()
	var final BingoBoard
	var finalNum int
OUTER:
	for _, num := range nums {
		for _, board := range boards {
			board.Mark(num)
			if board.CheckBingo() {
				final = board
				finalNum = num
				break OUTER
			}
		}
	}
	fmt.Println(final)
	fmt.Println(final.SumUnmarked(), finalNum, final.SumUnmarked()*finalNum)
}

type BingoBoard [][]int

func (b BingoBoard) SumUnmarked() int {
	sum := 0
	for i := range b {
		for j := range b[i] {
			if b[i][j] != -1 {
				sum += b[i][j]
			}
		}
	}
	return sum
}

func (b BingoBoard) CheckBingo() bool {
	return b.checkRows() || b.checkCols()
}

func (b BingoBoard) checkRows() bool {
	for i := range b {
		sum := 0
		for j := range b[i] {
			sum += b[i][j]
		}
		if sum == -5 {
			return true
		}
	}
	return false
}

func (b BingoBoard) checkCols() bool {
	for i := range b[0] {
		sum := 0
		for j := range b {
			sum += b[j][i]
		}
		if sum == -5 {
			return true
		}
	}
	return false
}

func (b BingoBoard) Mark(num int) {
	for i := range b {
		for j := range b[i] {
			if b[i][j] == num {
				b[i][j] = -1
			}
		}
	}
}

func ReadInput() (inputs []int, boards []BingoBoard) {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	scanner.Scan()
	inputs = breakNums(scanner.Text(), ",")

	// parse into bingo boards
	boards = []BingoBoard{}
	var board BingoBoard
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) < 5 {
			continue
		}
		if len(board) == 5 {
			boards = append(boards, board)
			board = BingoBoard{}
		}
		nums := breakNums(text, " ")
		board = append(board, nums)
	}
	return
}

func breakNums(str string, splitter string) []int {
	numStrs := strings.Split(str, splitter)
	nums := make([]int, 0, len(numStrs))
	for _, s := range numStrs {
		str := strings.TrimSpace(s)
		if str == "" {
			continue
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	return nums
}
