package day15

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
		out, err := part1(input, 2000000)
		if err != nil {
			return errors.Wrap(err, "part1")
		}
		fmt.Println("part1:", out)
		return nil
	},
}

func part1(input string, y int) (int, error) {
	pairs, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}

	yLine := map[Point]int{}

	for _, pair := range pairs {
		mRange := manhattanDistance(pair.Sensor, pair.Beacon)

		if y >= pair.Sensor.Y-mRange && y <= pair.Sensor.Y+mRange {
			left := mRange - abs(pair.Sensor.Y-y)
			count := 2*left + 1
			for i := 0; i < count; i++ {
				p := Point{pair.Sensor.X - left + i, y}
				if p == pair.Beacon {
					continue
				}
				yLine[p] = 1
			}
		}
	}
	return len(yLine), nil
}
