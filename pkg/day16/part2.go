package day16

import (
	"fmt"
	"sync"

	"github.com/ernestosuarez/itertools"
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
			return errors.Wrap(err, "part2")
		}
		fmt.Println("part2:", out)
		return nil
	},
}

func part2(input string) (int, error) {
	graph, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}

	tunnels := []string{}
	for k, v := range graph.Graph {
		if v == 0 {
			continue
		}
		tunnels = append(tunnels, k)
	}

	var maxPressure int
	matrix, err := doDijkstra(graph)
	if err != nil {
		return 0, errors.Wrap(err, "doDijkstra")
	}
	mu := sync.Mutex{}
	for i := 1; i < len(tunnels)/2; i++ {
		for v := range itertools.CombinationsStr(tunnels, i) {
			maxPressure1 := solveRecursive(matrix, graph, 0, 0, 0, "AA", v, 26)
			maxPressure2 := solveRecursive(matrix, graph, 0, 0, 0, "AA", createOpposite(tunnels, v), 26)
			mu.Lock()
			if newMax := maxPressure1 + maxPressure2; newMax > maxPressure {
				maxPressure = newMax
			}
			mu.Unlock()
		}
	}
	return maxPressure, nil
}

func createOpposite(all []string, partial []string) []string {
	new := []string{}

outer:
	for _, v := range all {
		for _, w := range partial {
			if v == w {
				continue outer
			}
		}
		new = append(new, v)
	}

	return new
}
