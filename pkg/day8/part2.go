package day8

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part2(input)
		if err != nil {
			return errors.Wrap(err, "part2")
		}
		fmt.Println(out)
		return nil
	},
}

/*
To measure the viewing distance from a given tree,
look up, down, left, and right from that tree;
stop if you reach an edge or at the first tree that
is the same height or taller than the tree under
consideration. (If a tree is right on the edge,
at least one of its viewing distances will be zero.)
*/
func part2(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	var maxScore int
	for x := 0; x < matrix.Width; x++ {
		for y := 0; y < matrix.Height; y++ {
			score := scoreTree(x, y, matrix)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore, nil
}

func scoreTree(x, y int, matrix Matrix) int {
	var score = 1
	for _, direction := range directions {
		score *= scoreDirection(x, y, matrix, direction)
	}
	return score
}

func scoreDirection(x, y int, matrix Matrix, direction []int) int {
	var score int
	deltaX, deltaY := direction[0], direction[1]
	current := matrix.get(x, y)
	for {
		if !matrix.valid(x+deltaX, y+deltaY) {
			return score
		}
		score++
		next := matrix.get(x+deltaX, y+deltaY)
		if next >= current {
			return score
		}
		x, y = x+deltaX, y+deltaY
	}
}
