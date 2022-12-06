package day5

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var operationReg = regexp.MustCompile(`move ([0-9]*) from ([0-9]*) to ([0-9]*)`)

func parseInput(i string) (*Ship, error) {
	s := bufio.NewScanner(strings.NewReader(i))
	containers, err := parseContainersInput(s)
	if err != nil {
		return nil, errors.Wrap(err, "parseContainersInput")
	}
	operations, err := parseOperations(s)
	if err != nil {
		return nil, errors.Wrap(err, "parseOperations")
	}
	return &Ship{
		Containers: containers,
		Operations: operations,
	}, nil
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parseOperations(s *bufio.Scanner) ([]*Operation, error) {
	out := []*Operation{}
	for s.Scan() {
		line := s.Text()
		matches := operationReg.FindStringSubmatch(line)
		if len(matches) != 4 {
			return nil, errors.Errorf("could not match: %s", line)
		}
		out = append(out, &Operation{
			Num: mustParseInt(matches[1]),
			Src: mustParseInt(matches[2]),
			Dst: mustParseInt(matches[3]),
		})
	}
	return out, nil
}

func parseContainersInput(s *bufio.Scanner) ([]*Stack, error) {
	containersInput := []string{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		containersInput = append(containersInput, line)
	}
	lookup := map[int]int{}
	for i, r := range containersInput[len(containersInput)-1] {
		if r == ' ' {
			continue
		}
		lookup[i] = int(r-'0') - 1
	}
	out := make([]*Stack, len(lookup))
	for i := 0; i < len(out); i++ {
		out[i] = &Stack{}
	}
	for i := len(containersInput) - 2; i >= 0; i-- {
		line := containersInput[i]
		for idx, dst := range lookup {
			r := line[idx]
			if r != ' ' {
				out[dst].push(r)
			}
		}
	}
	return out, nil
}
