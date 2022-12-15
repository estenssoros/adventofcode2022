package day15

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
		out, err := part2(input, 4000000)
		if err != nil {
			return errors.Wrap(err, "part2")
		}
		fmt.Println("part2:", out)
		return nil
	},
}

func part2(input string, tunningFreq int) (int, error) {
	pairs, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	sensors := map[Point]int{}
	for _, pair := range pairs {
		sensors[pair.Sensor] = manhattanDistance(pair.Sensor, pair.Beacon)
	}
	for y := 0; y < tunningFreq+1; y++ {
		for x := 0; x < tunningFreq+1; x++ {
			var inside bool
			for sensor, mRange := range sensors {
				if manhattanDistance(sensor, Point{x, y}) < mRange {
					inside = true
					x = sensor.X + mRange - abs(sensor.Y-y)
					break
				}
			}
			if !inside {
				fmt.Printf("found %d %d: %d\n", x, y, x*tunningFreq+y)
				return x*tunningFreq + y, nil
			}
		}
	}

	return 0, nil
}
