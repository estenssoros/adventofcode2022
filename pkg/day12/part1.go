package day12

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
		fmt.Println("part 1:", out)
		return nil
	},
}

func part1(input string) (int, error) {
	board, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	start, err := board.start()
	if err != nil {
		return 0, errors.Wrap(err, "board.start")
	}

	end, err := board.end()
	if err != nil {
		return 0, errors.Wrap(err, "board.end")
	}

	graph := Graph{}

	for y, row := range board.Matrix {
		for x := range row {
			point := Point{x, y}
			vertex := map[Point]int{}
			for _, move := range board.nextMoves(point) {
				vertex[move] = 1
			}
			graph[point] = vertex
		}
	}
	path, err := graph.path(start, end)
	if err != nil {
		return 0, errors.Wrap(err, "graph.path")
	}
	if draw {
		board.Draw(path.Points)
	}
	return path.Cost, err
}
