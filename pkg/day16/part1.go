package day16

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
		fmt.Println("part1:", out)
		return nil
	},
}

func part1(input string) (int, error) {
	graph, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}

	tunnels := []string{}
	for k, v := range graph.Cost {
		if v == 0 {
			continue
		}
		tunnels = append(tunnels, k)
	}
	// matrix, err := doDijkstra(graph)
	// if err != nil {
	// 	return 0, errors.Wrap(err, "doDijkstra")
	// }
	return solveRecursive(graph, 0, 0, 0, "AA", tunnels, 30), nil

}

func solveRecursive(graph *Graph, currentTime, currentPressure, currentFlow int, position string, remaining []string, limit int) int {
	max := currentPressure + (limit-currentTime)*currentFlow
	for _, v := range remaining {
		distance := graph.bestDistance(position, v) + 1
		if currentTime+distance < limit {
			possibleScore := solveRecursive(
				graph,
				currentTime+distance,
				currentPressure+distance*currentFlow,
				currentFlow+graph.Cost[v],
				v,
				removeFromList(remaining, v),
				limit,
			)
			if possibleScore > max {
				max = possibleScore
			}
		}
	}
	return max
}

func removeFromList(in []string, v string) []string {
	new := []string{}
	for _, i := range in {
		if i != v {
			new = append(new, i)
		}
	}
	return new
}
