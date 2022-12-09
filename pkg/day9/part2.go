package day9

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
			return errors.Wrap(err, "part1")
		}
		fmt.Println(out)
		return nil

	},
}

func part2(input string) (int, error) {
	steps, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	visited := map[Point]struct{}{}
	rope := NewRope(10)
	for _, step := range steps {
		for i := 0; i < step.Num; i++ {
			rope.move(step.Dir)
			visited[rope.Tail()] = struct{}{}
		}
	}
	return len(visited), nil

}
