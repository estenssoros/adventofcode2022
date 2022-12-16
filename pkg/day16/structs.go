package day16

import "encoding/json"

type Graph struct {
	Graph   map[string]int
	Edges   map[string][]string
	Mapping map[string]int
}

func (g Graph) String() string {
	ju, err := json.MarshalIndent(g, "", " ")
	if err != nil {
		panic(err)
	}
	return string(ju)
}
