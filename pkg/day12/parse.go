package day12

import (
	"bufio"
	"strings"
)

func parseInput(input string) (*Board, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	b := &Board{}
	for s.Scan() {
		line := s.Text()
		row := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = line[i]
		}
		b.Matrix = append(b.Matrix, row)
	}
	b.Height = len(b.Matrix)
	b.Width = len(b.Matrix[0])
	return b, nil
}
