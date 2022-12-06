package day5

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part1() },
}

func part1() error {
	ship, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	for _, op := range ship.Operations {
		for j := 0; j < op.Num; j++ {
			move(ship.Containers[op.Src-1], ship.Containers[op.Dst-1])
		}
	}
	ship.PrintTopRow()
	return nil
}

func (s *Ship) PrintTopRow() {
	for _, container := range s.Containers {
		fmt.Print(string(container.peek()))
	}
	fmt.Println()
}

func (s *Ship) Print() {
	for i, container := range s.Containers {
		fmt.Println(i+1, container.String())
	}
}
func move(stack1, stack2 *Stack) {
	stack2.push(stack1.pop())
}
