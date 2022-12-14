package day14

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
		fmt.Println("part2:", out)
		return nil
	},
}

func part2(input string) (int, error) {
	grid, err := parseInputPart2(input)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	var sands int
	for {
		sands++
		err = grid.dropSand(500, 0)
		if err != nil {
			break
		}
		if grid.get(500, 0) == sand {
			break
		}
	}
	fmt.Println(grid)
	return sands, nil
}
