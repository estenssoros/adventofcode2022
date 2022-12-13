package day{{.Day}}

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part{{.Part}}Cmd = &cobra.Command{
	Use:     "part{{.Part}}",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part{{.Part}}(input)
		if err != nil {
			return errors.Wrap(err, "part{{.Part}}")
		}
		fmt.Println("part{{.Part}}:", out)
		return nil
	},
}

func part{{.Part}}(input string) (int, error) {
	return 0, nil
}
