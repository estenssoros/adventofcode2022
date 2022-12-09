package day9

import (
	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed "input.txt"
var input string

func init() {
	Cmd.AddCommand(
		part1Cmd,
		part2Cmd,
	)
}

var Cmd = &cobra.Command{
	Use:   "day9",
	Short: "",
}
