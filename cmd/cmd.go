package cmd

import (
	"github.com/spf13/cobra"
)

func init(){
	cmd.AddCommand(
		// add commands here
	)
}

var cmd = &cobra.Command{
	Use:     "adventofcode2022",
	Short:   "",
}

func Execute() error {
	return cmd.Execute()
}
