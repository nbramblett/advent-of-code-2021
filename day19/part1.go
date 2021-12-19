package day19

import (
	"log"
	"math"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/sets"
	"github.com/nbramblett/advent-of-code-2021/util"
)

func Solve1() {
	lines := util.ReadLines("day19/sampleinput.txt")
	scanners := ReadScanners(lines)
	log.Println(len(scanners))
	log.Println(UniqueRelativeToFirst(scanners[0], scanners[1]))

	finalResult := sets.Set[Coordinate]{}
	// All of scanner 0 is counted
	// finalResult.Add(scanners[0]...)
	// for i := 1; i < len(scanners); i++ {
	// 	u := UniqueRelativeToFirst(finalResult.ToSlice(), scanners[i])
	// 	finalResult.Add(u...)
	// }
	log.Println(len(finalResult))
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
		// Ignore beacons out of reach of the offset
		if math.Abs(float64(c2.x)) > 1000 || math.Abs(float64(c2.y)) > 1000 || math.Abs(float64(c2.z)) > 1000 {
			continue
		}
		s2 = append(s2, c2)
	}
	return s2
}

func UniqueRelativeToFirst(s1, s2 Scanner) Scanner {
	s1Set := sets.New(s1...)
	unique := sets.New(s2...)
	variousOrientationsOfS2 := make([]Scanner, 48)
	var delta Coordinate
	for _, c := range s2 {
		for i, c2 := range Rotations(Orientations(c)...) {
			variousOrientationsOfS2[i] = append(variousOrientationsOfS2[i], c2)
		}
	}
	for _, c := range s1 {
		for _, ss2 := range variousOrientationsOfS2 {
			log.Println(ss2)
			for _, c2 := range ss2 {
				uniqueish := sets.Set[Coordinate]{}
				comm := 0
				delta = Coordinate{c.x - c2.x, c.y - c2.y, c.z - c2.z}
				for _, cc := range ss2 {
					cc2 := cc.Translate(delta)
					if s1Set.Contains(cc2) {
						comm++
					}
					uniqueish.Add(cc2)
				}
				log.Println(uniqueish)
				if comm >= 12 {
					unique = uniqueish
				}
			}
		}
	}
	u := unique.ToSlice()
	for i := range u {
		u[i] = u[i].Translate(delta)
	}
	return u
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

func Orientations(c Coordinate) []Coordinate {
	return []Coordinate{
		// {c.x, c.y, c.z},
		// {-c.x, c.y, c.z},
		// {c.x, -c.y, c.z},
		// {c.x, c.y, -c.z},
		// {-c.x, -c.y, c.z},
		{-c.x, c.y, -c.z},
		// {c.x, -c.y, -c.z},
		// {-c.x, -c.y, -c.z},
	}
}

func Rotations(coords ...Coordinate) []Coordinate {
	cs := sets.Set[Coordinate]{}
	for _, c := range coords {
		cs.Add(Coordinate{c.x, c.y, c.z})
	}
	return cs.ToSlice()
}
