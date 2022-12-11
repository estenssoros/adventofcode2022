package day11

import (
	_ "embed"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(
		part1Cmd,
		part2Cmd,
	)
}

var Cmd = &cobra.Command{
	Use:   "day11",
	Short: "",
}

//go:embed input.txt
var input string
