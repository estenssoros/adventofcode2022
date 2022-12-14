package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestPart1(t *testing.T) {
	out, err := part1(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 24, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 93, out)
}
