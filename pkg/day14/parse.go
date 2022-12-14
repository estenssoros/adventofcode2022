package day14

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/pkg/errors"
)

func parseInput(input string) (Grid, error) {
	lines, err := parseLines(input)
	if err != nil {
		return Grid{}, errors.Wrap(err, "parseLines")
	}
	return newGrid(lines), nil
}

func parseInputPart2(input string) (Grid, error) {
	lines, err := parseLines(input)
	if err != nil {
		return Grid{}, errors.Wrap(err, "parseLines")
	}
	return newGridPart2(lines), nil
}

func parseLines(input string) ([][]Point, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	lines := [][]Point{}
	for s.Scan() {
		line, err := parseLine(s.Text())
		if err != nil {
			return nil, errors.Wrap(err, "parseLine")
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func parseLine(line string) ([]Point, error) {
	fields := strings.Split(line, " -> ")
	points := []Point{}
	for _, field := range fields {
		var p Point
		_, err := fmt.Sscanf(field, "%d,%d", &p.X, &p.Y)
		if err != nil {
			return nil, errors.Wrap(err, "fmt.Sscanf")
		}
		points = append(points, p)
	}
	return points, nil
}

func minXY(lines [][]Point) Point {
	x := math.MaxInt64
	for _, line := range lines {
		for _, p := range line {
			if p.X < x {
				x = p.X
			}
		}
	}
	return Point{x, 0}
}

func maxXY(lines [][]Point) Point {
	x, y := -math.MaxInt64, -math.MaxInt64
	for _, line := range lines {
		for _, p := range line {
			if p.X > x {
				x = p.X
			}
			if p.Y > y {
				y = p.Y
			}
		}
	}
	return Point{x, y}
}
