package day{{.Day}}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = ``

func TestPart1(t *testing.T) {
	out, err := part1(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, out)
}
