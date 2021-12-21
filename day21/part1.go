package day21

import "log"

func Solve1() {
	player1, player2 := 2, 1 //my starting positions
	die := &Die{1}

	score1, score2 := 0, 0
	turn := 0
	for score1 < 1000 && score2 < 1000 {
		if turn%2 == 0 {
			score1 += Turn(&player1, die)
		} else {
			score2 += Turn(&player2, die)
		}
		turn += 3
	}
	losingScore := score1
	if score2 < score1 {
		losingScore = score2
	}
	log.Println(score1, score2, turn)
	log.Println(losingScore * turn)
}

func Turn(player *int, die *Die) int {
	for i := 0; i < 3; i++ {
		*player += die.Roll()
		*player %= 10
		if *player == 0 {
			*player = 10
		}
	}
	return *player
}

type Die struct {
	roll int
}

func (d *Die) Roll() int {
	k := d.roll
	d.roll++
	if d.roll == 101 {
		d.roll = 1
	}
	return k
}
