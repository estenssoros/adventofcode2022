package day15

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	Sensor  byte = 'S'
	Beacon  byte = 'B'
	Empty   byte = '.'
	Blocked byte = '#'
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Pair struct {
	Sensor Point
	Beacon Point
}

type Grid struct {
	min    Point
	max    Point
	matrix map[Point]byte
}

func (g Grid) String() string {
	b := strings.Builder{}
	for y := g.min.Y; y < g.max.Y; y++ {
		for x := g.min.X; x < g.max.X; x++ {
			b.WriteByte(g.get(x, y))
		}
		b.WriteByte('\n')
	}
	b.WriteString(fmt.Sprintf("min: %s\n", g.min))
	b.WriteString(fmt.Sprintf("max: %s\n", g.max))
	return b.String()
}

func newGrid(pairs []Pair) Grid {
	g := Grid{
		min:    Point{math.MaxInt, math.MaxInt},
		max:    Point{-math.MaxInt, -math.MaxInt},
		matrix: map[Point]byte{},
	}
	for _, pair := range pairs {
		g.SetPoint(pair.Beacon, Beacon)
		g.SetPoint(pair.Sensor, Sensor)
	}
	return g
}

func (g *Grid) SetPoint(p Point, b byte) {
	g.matrix[p] = b
	g.min = minPoint(g.min, p)
	g.max = maxPoint(g.max, p)
}

func (g Grid) get(x, y int) byte {
	return g.getPoint(Point{x, y})
}

func (g Grid) getPoint(p Point) byte {
	b, ok := g.matrix[p]
	if !ok {
		return Empty
	}
	return b
}

type Queue []*Item

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].Distance < q[j].Distance
}

func (q *Queue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*q = old[:n-1]
	return item
}

func (q *Queue) Push(x any) {
	// n := len(*q)
	item := x.(*Item)
	*q = append(*q, item)
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

type Item struct {
	Point    Point
	Distance int
}

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func (g *Grid) drawSensorRange(start Point) error {
	fmt.Printf("drawing sensor range: %s\n", start)
	if g.getPoint(start) != Sensor {
		return errors.New("point is not a sensor")
	}
	queue := &Queue{}
	heap.Init(queue)
	for _, direction := range directions {
		p := Point{start.X + direction[0], start.Y + direction[1]}
		item := &Item{
			Point:    p,
			Distance: manhattanDistance(start, p),
		}
		heap.Push(queue, item)
	}
	// var found bool
	for queue.Len() > 0 {
		next := heap.Pop(queue).(*Item).Point
		val := g.getPoint(next)
		switch val {
		case Beacon:
			// found = true
			// continue
			return nil
		case Sensor, Blocked:
			continue
		case Empty:
			g.SetPoint(next, Blocked)
		}
		for _, direction := range directions {
			p := Point{next.X + direction[0], next.Y + direction[1]}
			item := &Item{
				Point:    p,
				Distance: manhattanDistance(start, p),
			}
			heap.Push(queue, item)
		}
	}
	return nil
}

func (g *Grid) drawPairRange(pair Pair) error {
	fmt.Printf("drawing pair range %s - %s\n", pair.Sensor, pair.Beacon)
	distance := manhattanDistance(pair.Sensor, pair.Beacon)
	for y := g.min.Y; y < g.max.Y; y++ {
		for x := g.min.X; x < g.max.X; x++ {
			p := Point{x, y}
			newDistance := manhattanDistance(pair.Sensor, p)
			if newDistance > distance {
				continue
			}
			val := g.getPoint(p)
			switch val {
			case Empty:
				g.SetPoint(p, Blocked)
			case Sensor, Beacon, Blocked:
			}
		}
	}
	return nil
}

func (g *Grid) drawPairRangeAtLine(pair Pair, y int) error {
	fmt.Printf("drawing pair  %s - %s at line %d\n", pair.Sensor, pair.Beacon, y)
	distance := manhattanDistance(pair.Sensor, pair.Beacon)
	xRange := g.max.X - g.min.X + 1
	fmt.Println("xRange:", xRange)
	for x := g.min.X; x < g.max.X; x++ {
		p := Point{x, y}
		newDistance := manhattanDistance(pair.Sensor, p)
		if newDistance > distance {
			continue
		}
		switch g.getPoint(p) {
		case Empty:
			g.SetPoint(p, Blocked)
		case Sensor, Beacon, Blocked:
		}
	}
	return nil
}

func (g Grid) Line(y int) []byte {
	out := []byte{}
	for x := g.min.X; x < g.max.X; x++ {
		out = append(out, g.get(x, y))
	}
	return out
}

func manhattanDistance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func minPoint(p1, p2 Point) Point {
	return Point{
		X: min(p1.X, p2.X),
		Y: min(p1.Y, p2.Y),
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxPoint(p1, p2 Point) Point {
	return Point{
		X: max(p1.X, p2.X),
		Y: max(p1.Y, p2.Y),
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
