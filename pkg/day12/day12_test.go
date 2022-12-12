package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestPart1(t *testing.T) {
	out, err := part1(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 31, out)
}
