package gohelp

import "fmt"

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
