package day9

import (
	"fmt"
	"math"
)

var directions = map[string][]int{
	"U": {0, 1},
	"D": {0, -1},
	"R": {1, 0},
	"L": {-1, 0},
}

type Step struct {
	Dir string
	Num int
}

func (s Step) String() string {
	return fmt.Sprintf("move %s %d steps", s.Dir, s.Num)
}

type Point struct {
	X int
	Y int
}

func (p *Point) CatchUp(other Point) {
	p.X += nextDirection(p.X, other.X)
	p.Y += nextDirection(p.Y, other.Y)
}

func (p *Point) move(dir string) {
	direction, ok := directions[dir]
	if !ok {
		panic("missing:" + dir)
	}
	deltaX, deltaY := direction[0], direction[1]
	p.X += deltaX
	p.Y += deltaY
}

func (p Point) delta(other Point) (int, int) {
	return p.X - other.X, p.Y - other.Y
}

func (p Point) distance(other Point) int {
	deltaX, deltaY := p.delta(other)
	d := math.Sqrt(float64(deltaX*deltaX) + float64(deltaY*deltaY))
	return int(d)
}

type Rope struct {
	Knots []Point
}

func NewRope(n int) Rope {
	knots := make([]Point, n)
	return Rope{knots}
}

func (r Rope) Tail() Point {
	return r.Knots[len(r.Knots)-1]
}

func (r Rope) move(dir string) {
	r.Knots[0].move(dir)
	for i := 1; i < len(r.Knots); i++ {
		deltaX, deltaY := r.Knots[i-1].delta(r.Knots[i])
		if abs(deltaX) > 1 || abs(deltaY) > 1 {
			r.Knots[i].CatchUp(r.Knots[i-1])
		}
	}
}

func nextDirection(from, to int) int {
	if to == from {
		return 0
	} else if to > from {
		return 1
	}
	return -1
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
