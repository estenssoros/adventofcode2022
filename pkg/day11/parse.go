package day11

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseInput(input string) ([]*Monkey, error) {
	monkeys := []*Monkey{}
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "Monkey") {
			monkey, err := parseNewMonkey(line, s)
			if err != nil {
				return nil, errors.Wrap(err, "parsenNewMonkey")
			}
			monkeys = append(monkeys, monkey)
		}
	}
	return monkeys, nil
}

func parseNewMonkey(line string, s *bufio.Scanner) (*Monkey, error) {
	monkey := &Monkey{}
	if _, err := fmt.Sscanf(line, "Monkey %d:", &monkey.ID); err != nil {
		return nil, errors.Wrap(err, "fmt.Sscanf")
	}
	line, err := nextLinePrefix(s, "Starting items: ")
	if err != nil {
		return nil, errors.Wrap(err, "nextLinePrefix")
	}
	startingItems, err := stringToSliceOfInt(line)
	if err != nil {
		return nil, errors.Wrap(err, "stringToSliceOfInt")
	}
	monkey.Items = startingItems
	line, err = nextLinePrefix(s, "Operation: ")
	if err != nil {
		return nil, errors.Wrap(err, "nextLinePrefix")
	}
	_, err = fmt.Sscanf(line, "new = %s %s %s", &monkey.Operation.A, &monkey.Operation.Op, &monkey.Operation.B)
	if err != nil {
		return nil, errors.Wrap(err, "fmt.Sscanf")
	}
	line, err = nextLinePrefix(s, "Test: ")
	if err != nil {
		return nil, errors.Wrap(err, "nextLinePrefix")
	}
	_, err = fmt.Sscanf(line, "divisible by %d", &monkey.Test.Divisible)
	if err != nil {
		return nil, errors.Wrap(err, "fmt.Sscanf")
	}
	line, err = nextLinePrefix(s, "If true: ")
	if err != nil {
		return nil, errors.Wrap(err, "nextLinePrefix")
	}
	_, err = fmt.Sscanf(line, "throw to monkey %d", &monkey.Test.TrueDst)
	if err != nil {
		return nil, errors.Wrap(err, "fmt.Sscanf")
	}
	line, err = nextLinePrefix(s, "If false: ")
	if err != nil {
		return nil, errors.Wrap(err, "nextLinePrefix")
	}
	_, err = fmt.Sscanf(line, "throw to monkey %d", &monkey.Test.FalseDst)
	if err != nil {
		return nil, errors.Wrap(err, "fmt.Sscanf")
	}
	return monkey, nil
}

func nextLinePrefix(s *bufio.Scanner, prefix string) (string, error) {
	s.Scan()
	line := strings.TrimSpace(s.Text())
	if !strings.HasPrefix(line, prefix) {
		return "", errors.Errorf("could not parse [%s]", line)
	}
	return strings.TrimPrefix(line, prefix), nil
}

func stringToSliceOfInt(s string) ([]int64, error) {
	fields := strings.Split(s, ", ")
	out := make([]int64, len(fields))
	for i, field := range fields {
		c, err := strconv.Atoi(field)
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse: %s", field)
		}
		out[i] = int64(c)
	}
	return out, nil
}
