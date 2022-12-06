package day6

import (
	"fmt"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part2() },
}

func part2() error {
	fmt.Println(startOfPackageMarker(input, 14))
	return nil
}
