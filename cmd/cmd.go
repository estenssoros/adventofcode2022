package cmd

import (
	"github.com/estenssoros/adventofcode2022/pkg/day1"
	"github.com/estenssoros/adventofcode2022/pkg/day2"
	"github.com/estenssoros/adventofcode2022/pkg/day3"
	"github.com/estenssoros/adventofcode2022/pkg/day4"
	"github.com/estenssoros/adventofcode2022/pkg/day5"
	"github.com/estenssoros/adventofcode2022/pkg/day6"
	"github.com/estenssoros/adventofcode2022/pkg/day7"
	"github.com/estenssoros/adventofcode2022/pkg/day8"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(
		day1.Cmd,
		day2.Cmd,
		day3.Cmd,
		day4.Cmd,
		day5.Cmd,
		day6.Cmd,
		day7.Cmd,
		day8.Cmd,
	)
}

var cmd = &cobra.Command{
	Use:   "adventofcode2022",
	Short: "",
}

func Execute() error {
	return cmd.Execute()
}
