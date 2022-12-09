package day4

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
	pairs, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	var sum int
	for _, pair := range pairs {
		if pair.Section1.Overlaps(pair.Section2) || pair.Section2.Overlaps(pair.Section1) {
			sum++
		}
	}
	return sum, nil
}

func (s Section) Overlaps(other Section) bool {
	return s.Start >= other.Start && s.Start <= other.End
}
