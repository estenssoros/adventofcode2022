package day4

import (
	"bufio"
	"strconv"
	"strings"

	_ "embed"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

//go:embed input.txt
var input string

func init() {
	Cmd.AddCommand(
		part1Cmd,
		part2Cmd,
	)
}

var Cmd = &cobra.Command{
	Use:   "day4",
	Short: "",
}

type Pair struct {
	Section1 Section
	Section2 Section
}

type Section struct {
	Start int
	End   int
}

func parseInput(input string) ([]Pair, error) {
	pairs := []Pair{}
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		sections := strings.Split(line, ",")
		if len(sections) != 2 {
			return nil, errors.Errorf("could not parse: %s", line)
		}
		section1, err := parseSection(sections[0])
		if err != nil {
			return nil, errors.Wrap(err, "parseSection")
		}
		section2, err := parseSection(sections[1])
		if err != nil {
			return nil, errors.Wrap(err, "parseSection")
		}
		pairs = append(pairs, Pair{section1, section2})
	}
	return pairs, nil
}

func parseSection(s string) (Section, error) {
	fields := strings.Split(s, "-")
	if len(fields) != 2 {
		return Section{}, errors.Errorf("could not parse: %s", s)
	}
	start, err := strconv.Atoi(fields[0])
	if err != nil {
		return Section{}, errors.Wrap(err, "strconv.Atoi")
	}
	end, err := strconv.Atoi(fields[1])
	if err != nil {
		return Section{}, errors.Wrap(err, "strconv.Atoi")
	}
	return Section{start, end}, nil

}
