package day19

import (
	"log"
	"math"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/sets"
	"github.com/nbramblett/advent-of-code-2021/slices"
	"github.com/nbramblett/advent-of-code-2021/util"
)

var acceptableLevels = 12

func Solve1() {
	lines := util.ReadLines("day19/input.txt")
	scanners := ReadScanners(lines)
	log.Println(UniqueRelativeToFirst(scanners[0], scanners[1]))

	finalResult := Evaluate(scanners)
	log.Println(len(finalResult))
}

func Evaluate(scanners []Scanner) []Coordinate {
	finalResult := sets.Set[Coordinate]{}
	// All of scanner 0 is counted
	finalResult.Add(scanners[0]...)
	reevaluate := slices.Queue[int]{}
	for i := 1; i < len(scanners); i++ {
		uniqueSet := sets.New(scanners[i]...)
		foundMatch := false
		for j := 0; j < i; j++ {
			if slices.Contains(reevaluate, j) {
				continue
			}
			//log.Println("comparing", j, i)
			uniques, _ := UniqueRelativeToFirst(scanners[j], scanners[i])
			if uniques != nil {
				foundMatch = true
				us := sets.New(uniques...)
				//log.Println("matched!", us)
				uniqueSet = sets.Intersection(uniqueSet, us)
			}
		}
		if !foundMatch {
			//.Println("failed to match", i)
			reevaluate.Push(i)
			continue
		}
		//log.Println("unique count", i, uniqueSet.Len())
		finalResult.Add(uniqueSet.ToSlice()...)
	}
	bounce := 0
	for len(reevaluate) != 0 {
		s := reevaluate.Pop()
		//log.Println("retrying", s)
		uniqueSet := sets.New(scanners[s]...)
		foundMatch := false
		for i := 0; i < len(scanners); i++ {
			if slices.Contains(reevaluate, i) || i == s {
				continue
			}
			uniques, _ := UniqueRelativeToFirst(scanners[i], scanners[s])
			if uniques != nil {
				foundMatch = true
				us := sets.New(uniques...)
				uniqueSet = sets.Intersection(uniqueSet, us)
			}
		}
		if !foundMatch {
			log.Println("failed to match again", s)
			bounce++
			reevaluate.Push(s)
			if len(reevaluate) < bounce {
				panic("breaking infinite failure loop")
			}
			continue
		}
		bounce = 0
		//log.Println("unique count", s, uniqueSet.Len())
		finalResult.Add(uniqueSet.ToSlice()...)
		log.Println(len(reevaluate))
	}
	return finalResult.ToSlice()
}

func ReadScanners(lines []string) []Scanner {
	s := []Scanner{}
	var currentS Scanner
	for _, l := range lines {
		if l == "" {
			s = append(s, currentS)
			currentS = nil
			continue
		}
		if strings.HasPrefix(l, "---") {
			currentS = Scanner{}
			continue
		}
		i := util.StringsToInts(strings.Split(l, ","))
		if len(i) != 3 {
			panic(l)
		}
		currentS = append(currentS, Coordinate{i[0], i[1], i[2]})
	}
	return s
}

type Scanner []Coordinate

func Offset(s Scanner, c Coordinate) Scanner {
	s2 := Scanner{}
	for _, c1 := range s {
		c2 := c1.Translate(Coordinate{-c.x, -c.y, -c.z})
		s2 = append(s2, c2)
	}
	return s2
}

func UniqueRelativeToFirst(s1, s2 Scanner) (Scanner, Coordinate) {
	s1Set := sets.New(s1...)
	unique := sets.New(s2...)
	variousOrientationsOfS2 := RotateScanner(s2)
	foundMatching := false
	var delta Coordinate
OUTER:
	for _, c := range s1 {
		for _, ss2 := range variousOrientationsOfS2 {
			for _, c2 := range ss2 {
				uniqueish := sets.Set[Coordinate]{}
				comm := 0
				delta = Coordinate{c.x - c2.x, c.y - c2.y, c.z - c2.z}
				for i, cc := range ss2 {
					cc2 := cc.Translate(delta)
					if s1Set.Contains(cc2) {
						comm++
					} else {
						uniqueish.Add(s2[i])
					}
				}
				if comm >= acceptableLevels {
					foundMatching = true
					unique = uniqueish
					break OUTER
				}
			}
		}
	}
	if !foundMatching {
		return nil, delta
	}
	return unique.ToSlice(), delta
}

type Coordinate struct {
	x, y, z int
}

func (c Coordinate) Distance(c2 Coordinate) float64 {
	return math.Sqrt(float64((c.x-c2.x)*(c.x-c2.x) + (c.y-c2.y)*(c.y-c2.y) + (c.z-c2.z)*(c.z-c2.z)))
}

func (c Coordinate) Translate(motion Coordinate) Coordinate {
	return Coordinate{c.x + motion.x, c.y + motion.y, c.z + motion.z}
}

func Rotations(c Coordinate) []Coordinate {
	return []Coordinate{
		{c.x, c.y, c.z},
		{c.x, c.z, -c.y},
		{c.x, -c.y, -c.z},
		{c.x, -c.z, c.y},
		{-c.x, -c.y, c.z},
		{-c.x, c.z, c.y},
		{-c.x, c.y, -c.z},
		{-c.x, -c.z, c.y},

		{c.y, -c.x, c.z},
		{c.y, c.z, c.x},
		{c.y, c.x, -c.z},
		{c.y, -c.z, -c.x},
		{-c.y, c.x, c.z},
		{-c.y, c.z, -c.x},
		{-c.y, -c.x, -c.z},
		{-c.y, -c.z, c.x},

		{c.z, -c.y, c.x},
		{c.z, c.x, c.y},
		{c.z, c.y, -c.x},
		{c.z, -c.x, -c.y},
		{-c.z, c.y, c.x},
		{-c.z, c.x, -c.y},
		{-c.z, -c.y, -c.x},
		{-c.z, -c.x, c.y},
	}
}

func RotateScanner(s Scanner) []Scanner {
	ss := make([]Scanner, 24)
	for _, c := range s {
		for i, r := range Rotations(c) {
			ss[i] = append(ss[i], r)
		}
	}
	return ss
}
