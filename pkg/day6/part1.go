package day6

import (
	"fmt"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part1() },
}

func part1() error {
	fmt.Println(startOfPackageMarker(input, 4))
	return nil
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
	// lookup := map[byte]int{}
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
