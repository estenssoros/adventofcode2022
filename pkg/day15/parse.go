package day15

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func parseInput(input string) ([]Pair, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	pairs := []Pair{}
	for s.Scan() {
		line := s.Text()
		pair := Pair{}
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &pair.Sensor.X, &pair.Sensor.Y, &pair.Beacon.X, &pair.Beacon.Y)
		if err != nil {
			return nil, errors.Wrap(err, "fmt.Sscanf")
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}
