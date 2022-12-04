package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverLaps(t *testing.T) {
	p := Pair{
		Section1: Section{5, 7},
		Section2: Section{7, 9},
	}
	assert.True(t, true, p.Section1.Overlaps(p.Section2))
}
