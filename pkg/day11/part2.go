package day11

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
		out, err := monkeyBusiness(input, 10000, worryPart2)
		if err != nil {
			return errors.Wrap(err, "part2")
		}
		fmt.Println(out)
		return nil
	},
}

func worryPart2(i int64) int64 {
	return i
}
