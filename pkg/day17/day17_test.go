package day17

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestPart1(t *testing.T) {
	out, err := part1(testInput, 2022)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3068, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(testInput, 1000000000000)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1514285714288, out)
}

func TestBinary(t *testing.T) {
	out, err := strconv.ParseInt("0011110", 2, 64)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}
