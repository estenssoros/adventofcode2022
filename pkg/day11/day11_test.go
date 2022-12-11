package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

func TestPart1(t *testing.T) {
	out, err := monkeyBusiness(testInput, 20, worryPart1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 10605, out)
}

func TestPart2(t *testing.T) {
	out, err := monkeyBusiness(testInput, 10000, worryPart2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2713310158, out)
}
