package day8

import (
	"log"
	"reflect"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/util"
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

func convertNumber(inputs []string, opts map[rune]util.RuneSet) int {
	place := 1000
	value := 0
	for _, input := range inputs {
		value += convertDigit(input, opts) * place
		place /= 10
	}
	return value
}

func convertDigit(input string, opts map[rune]util.RuneSet) int {
	output := []rune{}
	for _, char := range input {
		output = append(output, opts[char].ToSlice()...)
	}

	for str, v := range baseCodes {
		if reflect.DeepEqual(util.NewRuneSet([]rune(str)...), util.NewRuneSet(output...)) {
			return v
		}
	}
	log.Fatalf("YOU GOOFED UP, it wanted %s which isn't an output", string(output))

	return 10000
}

func deduceLines(inputs string) map[rune]util.RuneSet {
	options := map[rune]util.RuneSet{}
	fiveLengths := []string{}
	sixLengths := []string{}
	three := ""
	two := ""
	// Fill in blank options
	for _, cha := range "abcdefg" {
		if _, ok := options[cha]; !ok {
			options[cha] = util.NewRuneSet([]rune("abcdefg")...)
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
				options[cha] = util.Intersection(options[cha], util.NewRuneSet('c', 'f'))
			}
		case 3: // 7
			for _, cha := range number {
				three = number
				options[cha] = util.Intersection(options[cha], util.NewRuneSet('a', 'c', 'f'))
			}
		case 4: // 4
			for _, cha := range number {
				options[cha] = util.Intersection(options[cha], util.NewRuneSet('b', 'c', 'd', 'f'))
			}
		case 5:
			fiveLengths = append(fiveLengths, number)
		case 6:
			sixLengths = append(sixLengths, number)
		}
	}

	// if you do '7' - '1', the only bit left is the top bar. So finding S_7 - S_1 gives us the character that aligns with 'a'
	threeTwoDisjunct := util.Minus(util.NewRuneSet([]rune(three)...), util.NewRuneSet([]rune(two)...))
	options[threeTwoDisjunct.ToSlice()[0]] = util.NewRuneSet('a')
	// Handle 6-length first. d, e, f
	if len(sixLengths) > 0 {
		uniqueChars := util.NewRuneSet()
		for _, s := range sixLengths {
			for _, s2 := range sixLengths {
				uniqueChars = util.Union(uniqueChars, util.Disjunction(util.NewRuneSet([]rune(s2)...), util.NewRuneSet([]rune(s)...)))
			}
		}
		for char := range uniqueChars {
			options[char] = util.Intersection(options[char], util.NewRuneSet('d', 'e', 'c'))
		}
	}

	// Handle 5-lengths next - there's 3 sets of variant characters. Find the conflicting characters and remove options. Only e, f, b, g, and c are viable
	if len(fiveLengths) > 1 {
		uniqueChars := util.NewRuneSet()
		for _, s := range fiveLengths {
			// if we subtract one from the five-lengths and get 3 characters remaining (rather than 4), we found 5, and know that 'adg' is left
			if set := util.Minus(util.NewRuneSet([]rune(s)...), util.NewRuneSet([]rune(two)...)); len(set) == 3 {
				for char := range set {
					options[char] = util.Intersection(options[char], util.NewRuneSet('a', 'd', 'g'))
				}
			}
			for _, s2 := range fiveLengths {
				uniqueChars = util.Union(uniqueChars, util.Disjunction(util.NewRuneSet([]rune(s2)...), util.NewRuneSet([]rune(s)...)))
			}
		}

		for char := range uniqueChars {
			options[char] = util.Intersection(options[char], util.NewRuneSet('b', 'c', 'e', 'f', 'g'))
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

func eliminateSingletons(options map[rune]util.RuneSet) {
	for cha, opts := range options {
		if len(opts) == 1 {
			recurse := false
			for c2, o2 := range options {
				if cha == c2 {
					continue
				}
				if o2.Contains(opts.ToSlice()[0]) && len(o2) > 1 {
					recurse = true
					options[c2] = util.Minus(o2, opts)
				}
			}
			if recurse {
				eliminateSingletons(options)
			}
		}
	}
}
