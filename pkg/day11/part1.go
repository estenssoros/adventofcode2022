package day11

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := monkeyBusiness(input, 20, worryPart1)
		if err != nil {
			return errors.Wrap(err, "part1")
		}
		fmt.Println(out)
		return nil
	},
}

func worryPart1(i int64) int64 {
	return i / 3
}

func monkeyBusiness(input string, rounds int, worry func(int64) int64) (int, error) {
	monkeys, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.Inspect(monkeys, worry)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].ItemsInspected > monkeys[j].ItemsInspected
	})

	return monkeys[0].ItemsInspected * monkeys[1].ItemsInspected, nil
}
