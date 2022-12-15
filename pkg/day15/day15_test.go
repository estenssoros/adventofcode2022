package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func TestPart1(t *testing.T) {
	out, err := part1(testInput, 10)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 26, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(testInput, 4000000)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, out)
}
