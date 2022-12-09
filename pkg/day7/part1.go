package day7

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

func part1(input string) (int, error) {
	dir, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}

	sizes := []int{}

	directories := []*Directory{}

	for _, c := range dir.Children {
		directories = append(directories, c)
	}

	for len(directories) > 0 {
		curr := directories[len(directories)-1]
		directories = directories[:len(directories)-1]
		if size := curr.Size(); size <= 100000 {
			sizes = append(sizes, size)
		}
		for _, c := range curr.Children {
			directories = append(directories, c)
		}
	}
	if size := dir.Size(); size <= 100000 {
		sizes = append(sizes, size)
	}

	return sum(sizes), nil
}

func sum(vals []int) int {
	var ttl int
	for _, v := range vals {
		ttl += v
	}
	return ttl
}
