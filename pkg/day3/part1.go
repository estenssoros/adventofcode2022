package day3

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part1() },
}

func part1() error {
	rucksacks, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	var score int
	for _, rucksack := range rucksacks {
		dup := rucksack.DuplicatedItem()
		score += scoreRune(dup)
	}
	fmt.Println(score)
	return nil
}

func isUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func scoreRune(r rune) int {
	if isUpperCase(r) {
		return int(r - 38)
	}
	return int(r - 96)
}
