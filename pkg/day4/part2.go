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
	RunE:    func(cmd *cobra.Command, args []string) error { return part2() },
}

func part2() error {
	pairs, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	var sum int
	for _, pair := range pairs {
		if pair.Section1.Overlaps(pair.Section2) || pair.Section2.Overlaps(pair.Section1) {
			sum++
		}
	}
	fmt.Println(sum)
	return nil
}

func (s Section) Overlaps(other Section) bool {
	return s.Start >= other.Start && s.Start <= other.End
}
