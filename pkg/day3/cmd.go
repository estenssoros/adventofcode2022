package day3

import (
	"bufio"
	_ "embed"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(
		part1Cmd,
		part2Cmd,
	)
}

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "",
}

//go:embed input.txt
var input string

type RuckSack struct {
	Compartment1 string
	Compartment2 string
}

func (r RuckSack) DuplicatedItem() rune {
	comp1 := map[rune]struct{}{}
	for _, r := range r.Compartment1 {
		comp1[r] = struct{}{}
	}
	for _, r := range r.Compartment2 {
		if _, ok := comp1[r]; ok {
			return r
		}
	}
	panic("could not find duplicated item")

}

func parseInput(input string) ([]RuckSack, error) {
	ruckSacks := []RuckSack{}
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		if len(line)%2 != 0 {
			return nil, errors.Errorf("could not parse: %s", line)
		}
		length := len(line)
		ruckSacks = append(ruckSacks, RuckSack{
			Compartment1: line[:length/2],
			Compartment2: line[length/2:],
		})

	}
	return ruckSacks, nil
}
