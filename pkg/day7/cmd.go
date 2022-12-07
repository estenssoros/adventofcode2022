package day7

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

//go:embed input.txt
var input string

var Cmd = &cobra.Command{
	Use:   "day7",
	Short: "",
}
