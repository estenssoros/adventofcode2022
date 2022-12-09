package day8

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part1(input)
		if err != nil {
			return errors.Wrap(err, "part1")
		}
		fmt.Println(out)
		return nil
	},
}

func part1(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	return countVisibleTrees(matrix), nil
}

var (
	directionRight = []int{1, 0}
	directionLeft  = []int{-1, 0}
	directionDown  = []int{0, 1}
	directionUp    = []int{0, -1}
)

var directions = [][]int{
	directionDown,
	directionLeft,
	directionRight,
	directionUp,
}

func countVisibleTrees(matrix Matrix) int {

	visible := make([][]bool, matrix.Height)
	for i := 0; i < matrix.Height; i++ {
		visible[i] = make([]bool, matrix.Width)
	}

	var count = matrix.Width*2 + (matrix.Width-2)*2

	for x := 1; x < matrix.Width-1; x++ {
		for y := 1; y < matrix.Height-1; y++ {
			if visible[y][x] {
				continue
			}
			for _, direction := range directions {
				if visibleInDirection(x, y, matrix, direction) && !visible[y][x] {
					count++
					visible[y][x] = true
				}
			}
		}
	}

	return count
}

func visibleInDirection(x, y int, matrix Matrix, direction []int) bool {
	deltaX, deltaY := direction[0], direction[1]
	current := matrix.get(x, y)
	for {
		if !matrix.valid(x+deltaX, y+deltaY) {
			return true
		}
		next := matrix.get(x+deltaX, y+deltaY)
		if next >= current {
			return false
		}
		x, y = x+deltaX, y+deltaY
	}
}
