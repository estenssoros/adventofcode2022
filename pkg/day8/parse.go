package day8

import (
	"bufio"
	"strings"
)

func parseInput(input string) (Matrix, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	out := [][]int{}
	var width int
	for s.Scan() {
		line := s.Text()
		if width == 0 {
			width = len(line)
		}
		row := make([]int, width)
		for i, r := range line {
			row[i] = int(r - '0')
		}
		out = append(out, row)
	}
	return Matrix{
		Cells:  out,
		Height: len(out),
		Width:  len(out[0]),
	}, nil
}
