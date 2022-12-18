package day17

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part1(input, 2022)
		if err != nil {
			return errors.Wrap(err, "part1")
		}
		fmt.Println("part1:", out)
		return nil
	},
}

func part1(input string, numIterations int) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rockChan := RockChan(ctx, Rocks)
	cave := Cave{
		wind: parseInput(ctx, input),
	}
	bar := progressbar.Default(int64(numIterations))
	for i := 0; i < numIterations; i++ {
		bar.Add(1)
		rock := <-rockChan
		cave.grow(3)
		cave.dropRock(rock)

	}
	fmt.Println(cave)
	return len(cave.grid), nil
}
