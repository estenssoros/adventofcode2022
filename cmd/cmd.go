package cmd

import (
	"github.com/estenssoros/adventofcode2022/pkg/day1"
	"github.com/estenssoros/adventofcode2022/pkg/day2"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(
		day1.Cmd,
		day2.Cmd,
	)
}

var cmd = &cobra.Command{
	Use:   "adventofcode2022",
	Short: "",
}

func Execute() error {
	return cmd.Execute()
}
