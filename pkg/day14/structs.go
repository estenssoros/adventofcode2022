package day14

import (
	"errors"
	"fmt"
	"strings"
)

var (
	sand  byte = 'o'
	rock  byte = '#'
	empty byte = '.'
)

type Grid struct {
	Min  Point
	Max  Point
	cave [][]byte
}

func (g Grid) String() string {
	b := strings.Builder{}
	for i := 0; i < len(g.cave); i++ {
		b.WriteString(string(g.cave[i]))
		b.WriteByte('\n')
	}
	b.WriteString("min: " + g.Min.String() + "\n")
	b.WriteString("max: " + g.Max.String() + "\n")
	b.WriteString(fmt.Sprintf("width: %d\n", g.width()))
	b.WriteString(fmt.Sprintf("depth: %d\n", g.depth()))
	return b.String()
}

func (g Grid) depth() int {
	return g.Max.Y - g.Min.Y + 1
}

func (g Grid) width() int {
	return g.Max.X - g.Min.X + 1
}

func newGrid(lines [][]Point) Grid {
	g := Grid{
		Min: minXY(lines),
		Max: maxXY(lines),
	}
	g.makeCave()
	for _, line := range lines {
		g.drawLine(line)
	}
	return g
}

func newGridPart2(lines [][]Point) Grid {
	min, max := minXY(lines), maxXY(lines)
	max.Y += 2
	min.X = 500 - max.Y
	max.X = 500 + max.Y
	g := Grid{
		Min: min,
		Max: max,
	}
	g.makeCave()
	for _, line := range lines {
		g.drawLine(line)
	}
	for i := 0; i < g.width(); i++ {
		g.cave[len(g.cave)-1][i] = rock
	}
	return g
}

func (g Grid) updatePoint(p Point) Point {
	return Point{
		X: p.X - g.Min.X,
		Y: p.Y,
	}
}

func (g *Grid) Set(x, y int, val byte) {
	x = x - g.Min.X
	g.cave[y][x] = val
}

func (g *Grid) SetPoint(p Point, val byte) {
	p = g.updatePoint(p)
	g.cave[p.Y][p.X] = val
}

func (g *Grid) drawLine(line []Point) {
	for i := 0; i < len(line)-1; i++ {
		g.drawPoints(line[i], line[i+1])
	}
}

func (g *Grid) drawPoints(p1, p2 Point) {
	g.SetPoint(p1, rock)
	g.SetPoint(p2, rock)
	for p1.X != p2.X || p1.Y != p2.Y {
		p1.X += mag(p2.X - p1.X)
		p1.Y += mag(p2.Y - p1.Y)
		g.SetPoint(p1, rock)
	}
}

func mag(i int) int {
	if i == 0 {
		return 0
	}
	if i > 0 {
		return 1
	}
	return -1
}

func (g *Grid) makeCave() {
	cave := make([][]byte, g.depth())
	for row := 0; row < g.depth(); row++ {
		cave[row] = make([]byte, g.width())
		for col := 0; col < g.width(); col++ {
			cave[row][col] = empty
		}
	}
	g.cave = cave
}

var ErrOutOfBounds = errors.New("out of bounds")

func (g *Grid) Valid(p Point) bool {
	p = g.updatePoint(p)
	if p.X < 0 || p.Y < 0 {
		return false
	}
	if p.X > g.width()-1 || p.Y > g.depth()-1 {
		return false
	}
	return true
}

func (g *Grid) dropSand(x, y int) error {
	p := Point{x, y}
	for {
		if !g.Valid(p.down()) {
			return ErrOutOfBounds
		}
		if g.down(p) != empty {
			break
		}
		p.Y++
	}

	if !g.Valid(p.leftDown()) {
		return ErrOutOfBounds
	}
	if g.leftDown(p) == empty {
		return g.dropSand(p.X-1, p.Y)
	}

	if !g.Valid(p.rightDown()) {
		return ErrOutOfBounds
	}
	if g.rightDown(p) == empty {
		return g.dropSand(p.X+1, p.Y)
	}
	g.SetPoint(p, sand)
	return nil
}

func (g *Grid) down(p Point) byte {
	p = g.updatePoint(p)
	return g.cave[p.Y+1][p.X]
}

func (g *Grid) get(x, y int) byte {
	p := g.updatePoint(Point{x, y})
	return g.cave[p.Y][p.X]
}

func (g *Grid) leftDown(p Point) byte {
	p = g.updatePoint(p)
	return g.cave[p.Y+1][p.X-1]
}

func (g *Grid) rightDown(p Point) byte {
	p = g.updatePoint(p)
	return g.cave[p.Y+1][p.X+1]
}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Point) down() Point {
	return Point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p Point) leftDown() Point {
	return Point{
		X: p.X - 1,
		Y: p.Y + 1,
	}
}

func (p Point) rightDown() Point {
	return Point{
		X: p.X + 1,
		Y: p.Y + 1,
	}
}
