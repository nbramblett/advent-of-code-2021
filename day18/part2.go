package day18

import "log"

func Solve2() {
	ns := ReadLines()
	maxMag := -1
	for i := 0; i < len(ns); i++ {
		for j := 0; j < len(ns); j++ {
			// Preserve the  i n t e g r i t y  of the pointers. Yes it's horribly wasteful, no I don't care.
			ns := ReadLines()
			if j == i {
				continue
			}
			n := Add(ns[i], ns[j])
			Resolve(n, true)
			m := n.Magnitude()
			if m > maxMag {
				maxMag = m
			}
		}
	}
	log.Println(maxMag)
}
