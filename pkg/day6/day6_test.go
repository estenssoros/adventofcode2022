package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTables = []struct {
	input string
	want  int
	k     int
}{
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 4},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6, 4},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 4},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 4},
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19, 14},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23, 14},
	{"nppdvjthqldpwncqszvftbrmjlhg", 23, 14},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29, 14},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26, 14},
}

func TestStartOfPacketMarker(t *testing.T) {
	for _, table := range testTables {
		t.Run(table.input, func(t *testing.T) {
			assert.Equal(t, table.want, startOfPackageMarker(table.input, table.k))
		})
	}
}

func TestCharAtIndex(t *testing.T) {
	assert.Equal(t, 0, charIndex('a'))
	assert.Equal(t, 25, charIndex('z'))
}
