package day16

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/RyanCarrier/dijkstra"
	"github.com/pkg/errors"
)

func parseInput(input string) (*Graph, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	edges := map[string][]string{}
	cost := map[string]int{}
	mapping := map[string]int{}
	var count int
	for s.Scan() {
		line := s.Text()
		part1, part2, err := splitString(line)
		if err != nil {
			return nil, errors.Wrap(err, "splitString")
		}
		valves := strings.Split(strings.TrimSpace(part2), ", ")
		var name string
		var rate int
		_, err = fmt.Sscanf(part1, "Valve %s has flow rate=%d;", &name, &rate)
		if err != nil {
			return nil, errors.Wrapf(err, "fmt.Sscanf")
		}
		edges[name] = valves
		cost[name] = rate
		mapping[name] = count
		count++
	}

	return &Graph{
		Cost:    cost,
		Edges:   edges,
		Mapping: mapping,
		Graph:   createDijkstraGraph(cost, mapping, edges),
		Cache:   map[string]map[string]int{},
	}, nil
}

func splitString(line string) (string, string, error) {
	idx := strings.Index(line, "valves ")
	if idx != -1 {
		idx += len("valves ")
		return line[:idx], line[idx:], nil
	}
	idx = strings.Index(line, "valve ")
	if idx == -1 {
		return "", "", errors.Errorf("could not parse: [%s]", line)
	}
	idx += len("valve ")
	return line[:idx], line[idx:], nil
}

func createDijkstraGraph(cost map[string]int, mapping map[string]int, edges map[string][]string) *dijkstra.Graph {
	dj := dijkstra.NewGraph()
	for k := range cost {
		dj.AddVertex(mapping[k])
	}
	for name, edges := range edges {
		for _, edge := range edges {
			dj.AddArc(mapping[name], mapping[edge], 1)
		}
	}
	return dj
}
