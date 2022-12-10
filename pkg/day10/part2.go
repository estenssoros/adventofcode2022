package day10

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

/*
the X register controls the horizontal position of a sprite
*/
func part2(input string) (CRT, error) {
	operations, err := parseInput(input)
	if err != nil {
		return CRT{}, errors.Wrap(err, "parseInput")
	}
	crt := newCRT(40, 6)
	crt.runOperations(operations)
	return crt, nil
}
