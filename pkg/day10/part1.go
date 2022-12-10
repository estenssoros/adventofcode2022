package day10

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

var samples = map[int]struct{}{
	20:  {},
	60:  {},
	100: {},
	140: {},
	180: {},
	220: {},
}

/*
noop
addx 3
addx -5

cycle command	x
1     noop      1
2     addx 3    1
3     -         4
4     addx -5   4
5     -         4
6     -         -1
*/

func part1(input string) (int, error) {
	operations, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}

	ch := make(chan Operation)
	out := processQueue(ch)
	for _, op := range operations {
		if !op.IsNoop {
			ch <- Operation{}
			ch <- op
		} else {
			ch <- op
		}
	}
	close(ch)

	return <-out, nil

}

func processQueue(operations <-chan Operation) chan int {
	ch := make(chan int)
	var x, cycle = 1, 1
	var out int
	go func() {
		defer close(ch)
		for op := range operations {
			x += op.AddX
			cycle++
			if _, ok := samples[cycle]; ok {
				out += cycle * x
			}
		}
		ch <- out
	}()
	return ch
}
