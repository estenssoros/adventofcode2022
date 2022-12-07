package day7

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part2(input) },
}

func part2(input string) error {
	totalSize := 70000000
	requiredSpace := 30000000
	dir, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	currentSpace := totalSize - dir.Size()
	toDelete := requiredSpace - currentSpace
	var minSize = int(math.MaxInt)
	directories := []*Directory{}
	for _, c := range dir.Children {
		directories = append(directories, c)
	}
	for len(directories) > 0 {
		curr := directories[len(directories)-1]
		directories = directories[:len(directories)-1]
		if size := curr.Size(); size >= toDelete {
			if size < minSize {
				minSize = size
			}
		}
		for _, c := range curr.Children {
			directories = append(directories, c)
		}
	}
	fmt.Println(minSize)
	return nil
}
