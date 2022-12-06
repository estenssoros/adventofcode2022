package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
move 10 from 7 to 6`

func TestPart1(t *testing.T) {
	ship, err := parseInput(testInput)
	if err != nil {
		t.Fatal(err)
	}
	wantOperation := []*Operation{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
		{10, 7, 6},
	}
	assert.Equal(t, wantOperation, ship.Operations)
	// ship.Part1()
	// for _, container := range ship.Containers {
	// 	fmt.Println(container.String())
	// }
	// wantContainers := []*Stack{
	// 	{'C'},
	// 	{'M'},
	// 	{'P', 'D', 'N', 'Z'},
	// }
	// assert.Equal(t, wantContainers, ship.Containers)
}
