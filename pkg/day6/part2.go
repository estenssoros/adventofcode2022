package day6

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

func part2(input string) (int, error) {
	return startOfPackageMarker(input, 14), nil
}
