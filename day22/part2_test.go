package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// I am not confident in my geometry lol

func TestIntersection(t *testing.T) {
	c1 := Cube{min: Point{0, 0, 0}, max: Point{2, 2, 2}}
	c2 := Cube{min: Point{1, 1, 1}, max: Point{3, 3, 3}}

	expected := Cube{min: Point{1, 1, 1}, max: Point{2, 2, 2}}
	assert.Equal(t, expected, Intersection(c1, c2))
}

func TestOverlaps(t *testing.T) {
	c1 := Cube{min: Point{0, 0, 0}, max: Point{2, 2, 2}}
	c2 := Cube{min: Point{1, 1, 1}, max: Point{3, 3, 3}}

	assert.True(t, Overlaps(c1, c2))

	// test "inside" check
	inside := c2.PointIn(c1.min) && c2.PointIn(c1.max)
	assert.False(t, inside)
	c1.min = Point{1, 1, 1}
	inside = c2.PointIn(c1.min) && c2.PointIn(c1.max)
	assert.True(t, inside)
}

func TestVolume(t *testing.T) {
	c := Cube{
		min:    Point{0, 0, 0},
		max:    Point{9, 9, 9},
		status: true,
		interiorExceptions: []*Cube{
			{min: Point{1, 1, 1}, max: Point{5, 5, 5}},
		},
	}

	assert.Equal(t, 10*10*10-5*5*5, c.Volume())

	c.interiorExceptions[0].interiorExceptions = []*Cube{
		{min: Point{2, 2, 2}, max: Point{3, 3, 3}},
	}

	assert.Equal(t, 10*10*10-5*5*5+8, c.Volume())
}
