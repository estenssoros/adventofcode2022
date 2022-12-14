package day14

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
		fmt.Println("part1:", out)
		return nil
	},
}

func part1(input string) (int, error) {
	grid, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	var sands int
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(grid)
			fmt.Println(sands)
			panic(r)
		}
	}()
	for {
		sands++
		err = grid.dropSand(500, 0)
		if err != nil {
			break
		}
	}
	return sands - 1, nil
}
