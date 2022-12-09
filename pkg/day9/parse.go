package day9

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func parseInput(input string) ([]Step, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	steps := []Step{}
	for s.Scan() {
		line := s.Text()
		step, err := parseStep(line)
		if err != nil {
			return nil, errors.Wrap(err, "parseStep")
		}
		steps = append(steps, step)
	}
	return steps, nil
}

func parseStep(line string) (Step, error) {
	step := Step{}
	_, err := fmt.Sscanf(line, "%s %d", &step.Dir, &step.Num)
	return step, err
}
