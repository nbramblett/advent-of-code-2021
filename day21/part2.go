package day21

import (
	"log"

	"github.com/nbramblett/advent-of-code-2021/util"
)

func Solve2() {

	startingState := BoardState{
		pos1: 2, pos2: 1, score1: 0, score2: 0, turn: 0,
	}
	cache := map[BoardState]Vector{}
	wins := Recurse(startingState, cache)

	log.Println(wins)
	log.Println(util.MinMax(wins.p1, wins.p2))
}

func Recurse(state BoardState, cache map[BoardState]Vector) Vector {
	if i, ok := cache[state]; ok {
		return i
	}
	if state.score1 >= 21 {
		return Vector{1, 0}
	} else if state.score2 >= 21 {
		return Vector{0, 1}
	}

	totals := Vector{0, 0}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				roll := i + j + k
				newState := state
				if state.turn == 0 {
					newState.pos1 += roll
					newState.pos1 %= 10
					if newState.pos1 == 0 {
						newState.pos1 = 10
					}
					newState.score1 += newState.pos1
				} else {
					newState.pos2 += roll
					newState.pos2 %= 10
					if newState.pos2 == 0 {
						newState.pos2 = 10
					}
					newState.score2 += newState.pos2
				}
				newState.turn++
				newState.turn %= 2
				wins := Recurse(newState, cache)
				totals.p1 += wins.p1
				totals.p2 += wins.p2
			}
		}
	}
	cache[state] = totals
	return totals
}

type Vector struct {
	p1, p2 int
}

type BoardState struct {
	score1, score2, pos1, pos2, turn int
}
