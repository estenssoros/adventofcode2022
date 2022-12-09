package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestPart1(t *testing.T) {
	out, err := part1(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 13, out)
}

var testInput2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestPart2(t *testing.T) {
	out, err := part2(testInput2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 36, out)
}
