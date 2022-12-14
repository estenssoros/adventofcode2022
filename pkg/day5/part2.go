package day5

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

func part2(input string) (string, error) {
	ship, err := parseInput(input)
	if err != nil {
		return "", errors.Wrap(err, "parseInput")
	}
	for _, op := range ship.Operations {
		moveKeepOrder(ship.Containers[op.Src-1], ship.Containers[op.Dst-1], op.Num)
	}
	return ship.TopRow(), nil
}

func moveKeepOrder(stack1, stack2 *Stack, num int) {
	tmpStack := &Stack{}
	for i := 0; i < num; i++ {
		tmpStack.push(stack1.pop())
	}
	for i := 0; i < num; i++ {
		stack2.push(tmpStack.pop())
	}
}
