package day16

import (
	"encoding/json"
	"sync"

	"github.com/RyanCarrier/dijkstra"
)

type Graph struct {
	Cost    map[string]int
	Edges   map[string][]string
	Mapping map[string]int
	Cache   map[string]map[string]int
	*dijkstra.Graph
	mu sync.Mutex
}

func (g Graph) String() string {
	ju, err := json.MarshalIndent(g, "", " ")
	if err != nil {
		panic(err)
	}
	return string(ju)
}

func (g *Graph) bestDistance(t1, t2 string) int {
	cache, hasT1 := g.Cache[t1]
	if hasT1 {
		best, hast2 := cache[t2]
		if hast2 {
			return best
		}
		bestPath, err := g.Shortest(g.Mapping[t1], g.Mapping[t2])
		if err != nil {
			panic(err)
		}
		g.Cache[t1][t2] = int(bestPath.Distance)
		return int(bestPath.Distance)
	}
	cache = map[string]int{}
	best, err := g.Shortest(g.Mapping[t1], g.Mapping[t2])
	if err != nil {
		panic(err)
	}
	cache[t2] = int(best.Distance)
	g.Cache[t1] = cache
	return int(best.Distance)
}

func (g *Graph) FilterTunnelPressure(pressure int) []string {
	out := []string{}
	for t, c := range g.Cost {
		if c == pressure {
			continue
		}
		out = append(out, t)
	}
	return out
}
