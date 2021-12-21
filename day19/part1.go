package day19

import (
	"log"
	"math"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/sets"
	"github.com/nbramblett/advent-of-code-2021/util"
)

var acceptableLevels = 12

func Solve1() {
	lines := util.ReadLines("day19/input.txt")
	scanners := ReadScanners(lines)

	finalResult, deltas := Evaluate(scanners)
	maxDistance := 0
	for _, d1 := range deltas {
		for _, d2 := range deltas {
			dist := abs(d1.x-d2.x) + abs(d1.y-d2.y) + abs(d1.z-d2.z)
			if dist > maxDistance {
				maxDistance = dist
			}
		}
	}
	log.Println(len(finalResult))
	log.Println(maxDistance)
}

func Evaluate(scanners []Scanner) ([]Coordinate, []Coordinate) {
	finalResult := sets.Set[Coordinate]{}
	// All of scanner 0 is counted
	finalResult.Add(scanners[0]...)
	deltas := []Coordinate{}
	convertedScanners := []Scanner{scanners[0]}
	for i := 1; i < len(scanners); i++ {
		log.Println(len(convertedScanners))
		uniqueSet := sets.New[Coordinate]()
		jlen := len(convertedScanners)
		matchedAny := false
		for j := 0; j < jlen; j++ {
			//log.Println("comparing", j, i)
			uniques, total, delta := UniqueRelativeToFirst(convertedScanners[j], scanners[i])
			if uniques != nil {
				us := sets.New(uniques...)
				//log.Println("matched!", us)
				if uniqueSet.Len() == 0 {
					uniqueSet = us
					convertedScanners = append(convertedScanners, total)
					matchedAny = true
					deltas = append(deltas, delta)
				} else {
					uniqueSet = sets.Intersection(uniqueSet, us)
				}
			}
		}
		if !matchedAny {
			scanners = append(scanners, scanners[i])
			continue
		}
		//log.Println("unique count", i, uniqueSet.Len())
		finalResult.Add(uniqueSet.ToSlice()...)
	}
	return finalResult.ToSlice(), deltas
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

func UniqueRelativeToFirst(s1, s2 Scanner) (Scanner, Scanner, Coordinate) {
	s1Set := sets.New(s1...)
	unique := sets.New(s2...)
	variousOrientationsOfS2 := RotateScanner(s2)
	foundMatching := false
	var delta Coordinate
	var convertedS2 sets.Set[Coordinate]
OUTER:
	for _, c := range s1 {
		for _, ss2 := range variousOrientationsOfS2 {
			for _, c2 := range ss2 {
				uniqueish := sets.Set[Coordinate]{}
				totalish := sets.Set[Coordinate]{}
				comm := 0
				delta = Coordinate{c.x - c2.x, c.y - c2.y, c.z - c2.z}
				for _, cc := range ss2 {
					cc2 := cc.Translate(delta)
					totalish.Add(cc2)
					if s1Set.Contains(cc2) {
						comm++
					} else {
						uniqueish.Add(cc2)
					}
				}
				if comm >= acceptableLevels {
					foundMatching = true
					unique = uniqueish
					convertedS2 = totalish
					break OUTER
				}
			}
		}
	}
	if !foundMatching {
		return nil, nil, Coordinate{}
	}
	return unique.ToSlice(), convertedS2.ToSlice(), delta
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
		{-c.x, -c.z, -c.y},

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

func abs(i int) int {
	return int(math.Abs(float64(i)))
}
