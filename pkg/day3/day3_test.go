package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreRune(t *testing.T) {
	assert.Equal(t, 16, scoreRune('p'))
	assert.Equal(t, 38, scoreRune('L'))
	assert.Equal(t, 42, scoreRune('P'))
	assert.Equal(t, 22, scoreRune('v'))
	assert.Equal(t, 20, scoreRune('t'))
	assert.Equal(t, 19, scoreRune('s'))
}

func TestIsUpperCase(t *testing.T) {
	assert.True(t, isUpperCase('P'))
	assert.False(t, isUpperCase('a'))
	assert.True(t, isUpperCase('B'))
}
