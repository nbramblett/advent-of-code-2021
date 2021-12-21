package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	cs := Rotations(Coordinate{1, 2, 3})
	assert.Len(t, cs, 24)
	for i := 0; i < 24; i++ {
		for j := 0; j < 24; j++ {
			if i != j {
				assert.NotEqual(t, Coordinate{0, 0, 0}, cs[j])
				assert.NotEqual(t, cs[i], cs[j])
			}
		}
	}
}
func TestRotateScanner(t *testing.T) {
	testScanner := Scanner{{1, 2, 3}, {2, 3, 4}}

	scanners := RotateScanner(testScanner)

	assert.Equal(t, testScanner[0], scanners[0][0], "Did not keep base in place")

	assert.Equal(t, Coordinate{1, -2, -3}, scanners[2][0], "Not consistent ordering")
	assert.Equal(t, Coordinate{2, -3, -4}, scanners[2][1], "Not consistent ordering")
}

func TestUniqueRelativeToFirst(t *testing.T) {
	acceptableLevels = 2
	scanners := []Scanner{
		{{1, 2, 3}, {2, 3, 4}, {1000, 2000, 3000}},
		{{3, 2, 5}, {4, 3, 6}, {0, 100, 200}},
		{{5, -3, 2}, {200, 0, 100}},
	}
	u, d := UniqueRelativeToFirst(scanners[0], scanners[1])
	assert.Equal(t, Coordinate{-2, 0, -2}, d)
	assert.Len(t, u, 1)
	assert.Equal(t, Coordinate{0, 100, 200}, u[0])
}

func TestEvaluate(t *testing.T) {
	acceptableLevels = 2
	scanners := []Scanner{
		{{1, 2, 3}, {2, 3, 4}, {1000, 2000, 3000}},
		{{3, 2, 5}, {4, 3, 6}, {0, 100, 200}},
		{{44, 33, 6}, {40, 130, 200}, {-123, -234, 435}},
	}

	finalResult := Evaluate(scanners)
	assert.ElementsMatch(t, []Coordinate{{1, 2, 3}, {2, 3, 4}, {0, 100, 200}, {1000, 2000, 3000}, {-123, -234, 435}}, finalResult)
}
