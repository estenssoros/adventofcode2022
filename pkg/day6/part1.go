package day6

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part1(input)
		if err != nil {
			return errors.Wrap(err, "part1")
		}
		fmt.Println(out)
		return nil
	},
}

func part1(input string) (int, error) {
	return startOfPackageMarker(input, 4), nil
}

/*
bvwbjplbgvbhsrlpgdmjqwftvncz
a b c d e f g h i j k l m n o p q r s t u v w x y z  right left

	1                                                  1
	1                                       1          2
	                                          1        3
	1                                       1 1        3     1
	                1                                  4     1
*/
func startOfPackageMarker(s string, k int) int {
	var left, right = 0, 0
	n := len(s)
	freq := make([]byte, 26)

	for right < n {

		freq[charIndex(s[right])]++

		for freq[charIndex(s[right])] > 1 {
			freq[charIndex(s[left])]--
			left += 1
		}
		if right-left+1 == k {
			return right + 1
		}
		right++
	}
	return 0
}

func charIndex(b byte) int {
	return int(b - 'a')
}
