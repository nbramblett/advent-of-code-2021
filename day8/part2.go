package day8

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

var baseCodes = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func Solve2() {
	lines := ReadInput()
	sum := 0
	for _, line := range lines {
		opts := deduceLines(line[0])
		vals := strings.Split(line[1], " ")
		sum += convertNumber(vals, opts)
	}
	log.Println(sum)
}

func convertNumber(inputs []string, opts map[rune]RuneSet) int {
	place := 1000
	value := 0
	for _, input := range inputs {
		value += convertDigit(input, opts) * place
		place /= 10
	}
	return value
}

func convertDigit(input string, opts map[rune]RuneSet) int {
	output := []rune{}
	for _, char := range input {
		output = append(output, opts[char].ToSlice()...)
	}

	for str, v := range baseCodes {
		if reflect.DeepEqual(NewRuneSet([]rune(str)...), NewRuneSet(output...)) {
			return v
		}
	}
	log.Fatalf("YOU GOOFED UP, it wanted %s which isn't an output", string(output))

	return 10000
}

func deduceLines(inputs string) map[rune]RuneSet {
	options := map[rune]RuneSet{}
	fiveLengths := []string{}
	sixLengths := []string{}
	three := ""
	two := ""
	// Fill in blank options
	for _, cha := range "abcdefg" {
		if _, ok := options[cha]; !ok {
			options[cha] = NewRuneSet([]rune("abcdefg")...)
		}
	}
	for _, number := range strings.Split(inputs, " ") {
		if len(number) == 7 {
			continue // 8 tells us nothing
		}

		switch len(number) {
		case 2: // 1
			for _, cha := range number {
				two = number
				options[cha] = Intersection(options[cha], NewRuneSet('c', 'f'))
			}
		case 3: // 7
			for _, cha := range number {
				three = number
				options[cha] = Intersection(options[cha], NewRuneSet('a', 'c', 'f'))
			}
		case 4: // 4
			for _, cha := range number {
				options[cha] = Intersection(options[cha], NewRuneSet('b', 'c', 'd', 'f'))
			}
		case 5:
			fiveLengths = append(fiveLengths, number)
		case 6:
			sixLengths = append(sixLengths, number)
		}
	}

	threeTwoDisjunct := Minus(NewRuneSet([]rune(three)...), NewRuneSet([]rune(two)...))
	options[threeTwoDisjunct.ToSlice()[0]] = NewRuneSet('a')
	// Handle 6-length first. d, e, f
	if len(sixLengths) > 0 {
		uniqueChars := NewRuneSet()
		for _, s := range sixLengths {
			for _, s2 := range sixLengths {
				uniqueChars = Union(uniqueChars, Disjunction(NewRuneSet([]rune(s2)...), NewRuneSet([]rune(s)...)))
			}
		}
		for char := range uniqueChars {
			options[char] = Intersection(options[char], NewRuneSet('d', 'e', 'c'))
		}
	}

	// Handle 5-lengths next - there's 3 sets of variant characters. Find the conflicting characters and remove options. Only e, f, b, g, and c are viable
	if len(fiveLengths) > 1 {
		uniqueChars := NewRuneSet()
		for _, s := range fiveLengths {
			// if we subtract one from the five-lengths and get 3 characters remaining (rather than 4), we found 5, and know that 'adg' is left
			if set := Minus(NewRuneSet([]rune(s)...), NewRuneSet([]rune(two)...)); len(set) == 3 {
				for char := range set {
					options[char] = Intersection(options[char], NewRuneSet('a', 'd', 'g'))
				}
			}
			for _, s2 := range fiveLengths {
				uniqueChars = Union(uniqueChars, Disjunction(NewRuneSet([]rune(s2)...), NewRuneSet([]rune(s)...)))
			}
		}

		for char := range uniqueChars {
			options[char] = Intersection(options[char], NewRuneSet('b', 'c', 'e', 'f', 'g'))
		}
	}

	eliminateSingletons(options)

	for _, set := range options {
		if len(set) > 1 {
			log.Fatalf("bad output %v", options)
		}
	}

	return options
}

func eliminateSingletons(options map[rune]RuneSet) {
	for cha, opts := range options {
		if len(opts) == 1 {
			recurse := false
			for c2, o2 := range options {
				if cha == c2 {
					continue
				}
				if o2.Contains(opts.ToSlice()[0]) && len(o2) > 1 {
					recurse = true
					options[c2] = Minus(o2, opts)
				}
			}
			if recurse {
				eliminateSingletons(options)
			}
		}
	}
}

type RuneSet map[rune]struct{}

func (s RuneSet) Add(r ...rune) {
	for _, ru := range r {
		s[ru] = struct{}{}
	}
}

func (s RuneSet) Remove(r ...rune) {
	for _, ru := range r {
		delete(s, ru)
	}
}

func (s RuneSet) Contains(r rune) bool {
	_, ok := s[r]
	return ok
}

func (s RuneSet) ToSlice() []rune {
	sl := make([]rune, 0, len(s))
	for key := range s {
		sl = append(sl, key)
	}
	return sl
}

func NewRuneSet(rs ...rune) RuneSet {
	newRuneSet := RuneSet{}
	for _, r := range rs {
		newRuneSet.Add(r)
	}
	return newRuneSet
}

func Union(rs ...RuneSet) RuneSet {
	newRuneSet := RuneSet{}
	for _, r := range rs {
		for key := range r {
			newRuneSet.Add(key)
		}
	}
	return newRuneSet
}

func Intersection(r1, r2 RuneSet) RuneSet {
	nr := RuneSet{}
	for key := range r1 {
		if _, ok := r2[key]; ok {
			nr.Add(key)
		}
	}
	return nr
}

func Minus(r1, r2 RuneSet) RuneSet {
	nr := RuneSet{}
	for key := range r1 {
		if _, ok := r2[key]; !ok {
			nr.Add(key)
		}
	}
	return nr
}

func Disjunction(r1, r2 RuneSet) RuneSet {
	nr1, nr2 := RuneSet{}, RuneSet{}
	for key := range r1 {
		if _, ok := r2[key]; !ok {
			nr1.Add(key)
		}
	}
	for key := range r2 {
		if _, ok := r1[key]; !ok {
			nr2.Add(key)
		}
	}
	return Union(nr1, nr2)
}

func (r RuneSet) String() string {
	s := "set["
	for key := range r {
		s = fmt.Sprintf("%s %c", s, key)
	}
	s = fmt.Sprintf("%s ]", s)
	return s
}
