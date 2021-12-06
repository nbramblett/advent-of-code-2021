package day4

import (
	"fmt"
)

func Solve2() {
	nums, boards := ReadInput()
	var final BingoBoard
	var finalNum int
	won := make([]bool, len(boards))
	for _, num := range nums {
	BOARD:
		for i, board := range boards {
			if won[i] {
				continue
			}
			board.Mark(num)
			if board.CheckBingo() {
				won[i] = true
				for _, wons := range won {
					if !wons {
						continue BOARD
					}
				}
				final = board
				finalNum = num
			}
		}
	}
	fmt.Println(final)
	fmt.Println(final.SumUnmarked(), finalNum, final.SumUnmarked()*finalNum)
}
