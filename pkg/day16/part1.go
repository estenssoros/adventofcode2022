package day16

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra"
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
	for k, v := range graph.Graph {
		if v == 0 {
			continue
		}
		tunnels = append(tunnels, k)
	}
	matrix, err := doDijkstra(graph)
	if err != nil {
		return 0, errors.Wrap(err, "doDijkstra")
	}
	return solveRecursive(matrix, graph, 0, 0, 0, "AA", tunnels, 30), nil

}

func solveRecursive(matrix map[string]map[string]int, graph *Graph, currentTime, currentPressure, currentFlow int, position string, remaining []string, limit int) int {
	max := currentPressure + (limit-currentTime)*currentFlow
	for _, v := range remaining {
		distance := matrix[position][v] + 1
		if currentTime+distance < limit {
			possibleScore := solveRecursive(
				matrix,
				graph,
				currentTime+distance,
				currentPressure+distance*currentFlow,
				currentFlow+graph.Graph[v],
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

func doDijkstra(graph *Graph) (map[string]map[string]int, error) {
	dj := dijkstra.NewGraph()
	for k := range graph.Graph {
		dj.AddVertex(graph.Mapping[k])
	}
	for name, edges := range graph.Edges {
		for _, edge := range edges {
			dj.AddArc(graph.Mapping[name], graph.Mapping[edge], 1)
		}
	}
	matrix := map[string]map[string]int{}
	for name, idx := range graph.Mapping {
		matrix[name] = map[string]int{}
		for name2, idx2 := range graph.Mapping {
			if name == name2 {
				continue
			}
			best, err := dj.Shortest(idx, idx2)
			if err != nil {
				return nil, errors.Wrap(err, "dj.Shortest")
			}
			matrix[name][name2] = int(best.Distance)
		}
	}
	return matrix, nil
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
