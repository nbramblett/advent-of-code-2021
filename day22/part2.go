package day22

import (
	"log"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/util"
)

func Solve2() {
	lines := util.ReadLines("day22/sampleinput.txt")
	onOff, insts := SplitLines(lines)
	cubes := EvalCubes(onOff, insts)
	totalOn := 0
	overlaps := []Cube{}
	for i, c := range cubes {
		if c == nil {
			continue
		}
		for j, c2 := range c.interiorExceptions {
			if j <= i || c == nil {
				continue
			}
			if Overlaps(*c, *c2) {
				overlaps = append(overlaps, Intersection(*c, *c2))
			}
		}
	}
	for _, c := range cubes {
		v := c.Volume()
		log.Println("final volume", v)
		totalOn += v
	}
	log.Println(len(overlaps))
	for _, o := range overlaps {
		totalOn -= o.Volume()
	}
	log.Println(totalOn)
	log.Println(totalOn - 2758514936282235)
}

func EvalCubes(onOff, insts []string) []*Cube {
	cubes := []*Cube{}
	for i := range onOff {
		c := ParseCube(insts[i])
		c.status = onOff[i] == "on"
		for j, cube := range cubes {
			if cube == nil || !Overlaps(*c, *cube) {
				continue
			}
			// If a cube is totally "swallowed", just nil it out bc it doesn't matter anymore
			if c.PointIn(cube.min) && c.PointIn(cube.max) {
				cubes[j] = nil
				continue
			}
			CarveCube(c, cube)
		}
		if c.status {
			cubes = append(cubes, c)
		}
	}
	return cubes
}

type Cube struct {
	min, max           Point
	status             bool
	interiorExceptions []*Cube
}

func CarveCube(a, e *Cube) {
	if a.status == e.status {
		for i, ec := range e.interiorExceptions {
			if ec != nil && Overlaps(*a, *ec) {
				if a.PointIn(ec.min) && a.PointIn(ec.max) {
					e.interiorExceptions[i] = nil
				}
				CarveCube(a, ec)
			}
		}
		return
	}
	i := Intersection(*a, *e)
	log.Println("carving out", i.Volume())
	i.status = a.status
	e.interiorExceptions = append(e.interiorExceptions, &i)
}

// Volume returns the total number of cells in the cube that actually match the cube's status
func (c *Cube) Volume() int {
	if c == nil {
		return 0
	}
	// max is inclusive :(
	base := (c.max.x - c.min.x + 1) * (c.max.y - c.min.y + 1) * (c.max.z - c.min.z + 1)
	if len(c.interiorExceptions) == 0 {
		return base
	}
	overlaps := []Cube{}
	//Prevent double-counting overlap
	for i, ie := range c.interiorExceptions {
		if ie == nil {
			continue
		}
		for j, ie2 := range c.interiorExceptions {
			if j <= i || ie2 == nil {
				continue
			}
			if Overlaps(*ie, *ie2) {
				overlaps = append(overlaps, Intersection(*ie, *ie2))
			}
		}
	}
	for _, ie := range c.interiorExceptions {
		log.Println("subtracting", base, ie.Volume(), base-ie.Volume())
		base -= ie.Volume()
	}
	for _, o := range overlaps {
		log.Println("adding back", base, o.Volume(), base+o.Volume())
		base += o.Volume()
	}
	return base
}

func (c Cube) PointIn(p Point) bool {
	return p.x >= c.min.x && p.x <= c.max.x &&
		p.y >= c.min.y && p.y <= c.max.y &&
		p.z >= c.min.z && p.z <= c.max.z
}

func Overlaps(c1, c2 Cube) bool {
	return !(c2.min.x > c1.max.x || c2.min.y > c1.max.y || c2.min.z > c1.max.z ||
		c2.max.x < c1.min.x || c2.max.y < c1.min.y || c2.max.z < c1.min.z)
}

func Intersection(c1, c2 Cube) Cube {
	max, min := Point{}, Point{}
	max.x, _ = util.MinMax(c2.max.x, c1.max.x)
	max.y, _ = util.MinMax(c2.max.y, c1.max.y)
	max.z, _ = util.MinMax(c2.max.z, c1.max.z)
	_, min.x = util.MinMax(c2.min.x, c1.min.x)
	_, min.y = util.MinMax(c2.min.y, c1.min.y)
	_, min.z = util.MinMax(c2.min.z, c1.min.z)
	return Cube{min: min, max: max}
}

func ParseCube(l string) *Cube {
	coords := strings.Split(l, ",")
	x := util.StringsToInts(strings.Split(coords[0][2:], ".."))
	y := util.StringsToInts(strings.Split(coords[1][2:], ".."))
	z := util.StringsToInts(strings.Split(coords[2][2:], ".."))
	x[0], x[1] = util.MinMax(x[0], x[1])
	y[0], y[1] = util.MinMax(y[0], y[1])
	z[0], z[1] = util.MinMax(z[0], z[1])

	return &Cube{min: Point{x[0], y[0], z[0]}, max: Point{x[1], y[1], z[1]}}
}
