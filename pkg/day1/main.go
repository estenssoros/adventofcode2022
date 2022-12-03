package day1

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
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
	Use:   "day1",
	Short: "",
}

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part1() },
}

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part2() },
}

//go:embed input.txt
var input string

func part1() error {
	elves, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	max := maxElf(elves)
	fmt.Println(max)
	return nil
}

func part2() error {
	elves, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	var sum int
	for i := 0; i < 3; i++ {
		sum += elves[i]
	}
	fmt.Println(sum)
	return nil

}

func maxElf(elves []int) int {
	var max int
	for _, elf := range elves {
		if elf > max {
			max = elf
		}
	}
	return max
}

func parseInput(input string) ([]int, error) {
	elves := []int{}
	s := bufio.NewScanner(strings.NewReader(input))
	var calories int
	for s.Scan() {
		line := s.Text()
		if line == "" {
			elves = append(elves, calories)
			calories = 0
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.Wrap(err, "strconv.Atoi")
		}
		calories += i
	}
	if calories > 0 {
		elves = append(elves, calories)
	}
	return elves, nil
}
