package day9

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
	steps, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	visited := map[Point]struct{}{}
	head, tail := Point{}, Point{}
	for _, step := range steps {
		for i := 0; i < step.Num; i++ {
			headPrevious := head
			head.move(step.Dir)
			if head.distance(tail) > 1 {
				tail = headPrevious
			}
			visited[tail] = struct{}{}
		}
	}

	return len(visited), nil
}
