package day12

import (
	_ "embed"

	"github.com/spf13/cobra"
)

var draw bool

func init() {
	Cmd.AddCommand(
		part1Cmd,
		part2Cmd,
	)
	Cmd.PersistentFlags().BoolVarP(&draw, "draw", "", false, "draw solution board")
}

//go:embed input.txt
var input string

var Cmd = &cobra.Command{
	Use:   "day12",
	Short: "",
}
