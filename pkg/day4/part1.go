package day4

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
	pairs, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	var sum int
	for _, pair := range pairs {
		if pair.Section1.Contains(pair.Section2) || pair.Section2.Contains(pair.Section1) {
			sum++
		}
	}
	fmt.Println(sum)
	return nil
}

func (s Section) Contains(other Section) bool {
	return s.Start <= other.Start && s.End >= other.End
}
