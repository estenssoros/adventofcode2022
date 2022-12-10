package day10

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func parseInput(input string) ([]Operation, error) {
	operations := []Operation{}
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		if line == "noop" {
			operations = append(operations, Operation{IsNoop: true})
			continue
		}
		var addX int
		_, err := fmt.Sscanf(line, "addx %d", &addX)
		if err != nil {
			return nil, errors.Wrap(err, "fmt.Sscanf")
		}
		operations = append(operations, Operation{AddX: addX})

	}
	return operations, nil
}
