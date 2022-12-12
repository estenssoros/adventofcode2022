package day12

import (
	"fmt"
	"math"

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

func part2(input string) (int, error) {
	board, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	starts := board.starts()
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
	var (
		minDistance = math.MaxInt
		minPath     Path
	)
	for _, start := range starts {
		path, err := graph.path(start, end)
		if err != nil {
			if err == ErrNoPath {
				continue
			}
			return 0, errors.Wrap(err, "graph.path")
		}
		if path.Cost < minDistance {
			minDistance = path.Cost
			minPath = path
		}
	}
	if draw {
		board.Draw(minPath.Points)
	}
	return minDistance, nil
}
