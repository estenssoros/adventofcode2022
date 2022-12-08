package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `30373
25512
65332
33549
35390`

func TestCountVisibleTrees(t *testing.T) {
	matrix, err := parseInput(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 21, countVisibleTrees(matrix))
}

func TestVisibleInDirection(t *testing.T) {
	matrix, err := parseInput(testInput)
	if err != nil {
		t.Fatal(err)
	}
	//The top-left 5 is visible from the left and top
	assert.True(t, visibleInDirection(1, 1, matrix, directionUp))
	assert.True(t, visibleInDirection(1, 1, matrix, directionLeft))

	//The top-middle 5 is visible from the top and right
	assert.True(t, visibleInDirection(2, 1, matrix, directionUp))
	assert.True(t, visibleInDirection(2, 1, matrix, directionRight))

	//The top-right 1 is not visible from any direction
	assert.False(t, visibleInDirection(3, 1, matrix, directionDown))
	assert.False(t, visibleInDirection(3, 1, matrix, directionUp))
	assert.False(t, visibleInDirection(3, 1, matrix, directionLeft))
	assert.False(t, visibleInDirection(3, 1, matrix, directionRight))

	//The left-middle 5 is visible, but only from the right
	assert.False(t, visibleInDirection(1, 2, matrix, directionDown))
	assert.False(t, visibleInDirection(1, 2, matrix, directionUp))
	assert.False(t, visibleInDirection(1, 2, matrix, directionLeft))
	assert.True(t, visibleInDirection(1, 2, matrix, directionRight))

	//The center 3 is not visible from any direction
	assert.False(t, visibleInDirection(2, 2, matrix, directionDown))
	assert.False(t, visibleInDirection(2, 2, matrix, directionUp))
	assert.False(t, visibleInDirection(2, 2, matrix, directionLeft))
	assert.False(t, visibleInDirection(2, 2, matrix, directionRight))

	// The right-middle 3 is visible from the right
	assert.False(t, visibleInDirection(3, 2, matrix, directionDown))
	assert.False(t, visibleInDirection(3, 2, matrix, directionUp))
	assert.False(t, visibleInDirection(3, 2, matrix, directionLeft))
	assert.True(t, visibleInDirection(3, 2, matrix, directionRight))
}

func TestScoreInDirection(t *testing.T) {
	matrix, err := parseInput(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, scoreDirection(2, 1, matrix, directionUp))
	assert.Equal(t, 1, scoreDirection(2, 1, matrix, directionLeft))
	assert.Equal(t, 2, scoreDirection(2, 1, matrix, directionRight))
	assert.Equal(t, 2, scoreDirection(2, 1, matrix, directionDown))

	assert.Equal(t, 2, scoreDirection(2, 3, matrix, directionUp))
	assert.Equal(t, 2, scoreDirection(2, 3, matrix, directionLeft))
	assert.Equal(t, 2, scoreDirection(2, 3, matrix, directionRight))
	assert.Equal(t, 1, scoreDirection(2, 3, matrix, directionDown))

	assert.Equal(t, 0, scoreDirection(0, 0, matrix, directionUp))
	assert.Equal(t, 0, scoreDirection(0, 0, matrix, directionLeft))
}
