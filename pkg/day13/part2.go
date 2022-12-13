package day13

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part2(input)
		if err != nil {
			return errors.Wrap(err, "part1")
		}
		fmt.Println("part2:", out)
		return nil
	},
}

func part2(input string) (int, error) {
	packets, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}
	all := []any{}
	for _, packet := range packets {
		all = append(all, packet.l1, packet.l2)
	}
	var divider1 any
	if err := json.Unmarshal([]byte("[[2]]"), &divider1); err != nil {
		return 0, errors.Wrap(err, "json.Unmarshal")
	}
	var divider2 any
	if err := json.Unmarshal([]byte("[[6]]"), &divider2); err != nil {
		return 0, errors.Wrap(err, "json.Unmarshal")
	}
	all = append(all, divider1, divider2)
	sort.Slice(all, func(i, j int) bool {
		return isValid(all[i], all[j])
	})
	var out = 1
	for i, packet := range all {
		packetBytes, err := json.Marshal(packet)
		if err != nil {
			return 0, errors.Wrap(err, "json.Marshal")
		}
		switch string(packetBytes) {
		case "[[2]]", "[[6]]":
			out *= i + 1
		}
	}
	return out, nil
}
