package day3

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part2() },
}

func part2() error {
	groups, err := parseInputPart2(input)
	if err != nil {
		return errors.Wrap(err, "parseInputPart2")
	}
	var score int
	for _, group := range groups {
		badge, err := group.CommonLetter()
		if err != nil {
			return errors.Wrap(err, "group.CommonLetter")
		}
		score += scoreRune(badge)
	}
	fmt.Println(score)
	return nil
}

type Group struct {
	RuckSacks []string
}

func (g Group) CommonLetter() (rune, error) {
	lookups := g.Lookups()
	first := lookups[0]
	for _, l := range lookups[1:] {
		for r := range first {
			if _, ok := l[r]; !ok {
				delete(first, r)
			}
		}
	}
	if len(first) != 1 {
		return 'a', errors.Errorf("found duplicates: %v", first)
	}
	var r rune
	for key := range first {
		r = key
	}

	return r, nil
}

func (g Group) Lookups() []map[rune]struct{} {
	lookups := make([]map[rune]struct{}, len(g.RuckSacks))
	for i, rucksack := range g.RuckSacks {
		lookups[i] = makeLookup(rucksack)
	}
	return lookups
}

func makeLookup(s string) map[rune]struct{} {
	l := map[rune]struct{}{}
	for _, r := range s {
		l[r] = struct{}{}
	}
	return l
}

func parseInputPart2(input string) ([]Group, error) {
	groups := []Group{}
	s := bufio.NewScanner(strings.NewReader(input))
	group := []string{}
	for s.Scan() {
		group = append(group, s.Text())
		if len(group) == 3 {
			groups = append(groups, Group{group})
			group = nil
		}
	}
	return groups, nil
}
