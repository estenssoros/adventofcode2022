package day12

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type Point struct {
	X int
	Y int
}

type Board struct {
	Height int
	Width  int
	Matrix [][]byte
}

func (b Board) String() string {
	builder := strings.Builder{}
	for i, row := range b.Matrix {
		builder.WriteString(string(row))
		if i < len(b.Matrix)-1 {
			builder.WriteByte('\n')
		}
	}
	return builder.String()
}

func (b Board) start() (Point, error) {
	return b.findChar('S')
}

func (b Board) starts() []Point {
	points := []Point{}
	for y, row := range b.Matrix {
		for x, el := range row {
			switch el {
			case 'a', 'S':
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func (b Board) end() (Point, error) {
	return b.findChar('E')
}

func (b Board) findChar(char byte) (Point, error) {
	for y, row := range b.Matrix {
		for x, el := range row {
			if el == char {
				return Point{x, y}, nil
			}
		}
	}
	return Point{}, errors.New("could not find char")
}

func (b Board) get(x, y int) byte {
	return b.Matrix[y][x]
}

func (b Board) valid(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= b.Width || y >= b.Height {
		return false
	}
	return true
}

func (b Board) nextMoves(p Point) []Point {
	current := b.get(p.X, p.Y)

	points := []Point{}

	for _, d := range directions {
		x, y := p.X+d[0], p.Y+d[1]

		if !b.valid(x, y) {
			continue
		}

		next := b.get(x, y)
		if current == 'S' {
			current = 'a'
		}
		if next == 'E' {
			next = 'z'
		}
		if int(next)-int(current) < 2 {
			points = append(points, Point{x, y})
		}
	}
	return points
}

func (b *Board) Draw(points []Point) {
	has := map[Point]struct{}{}
	for _, p := range points {
		has[p] = struct{}{}
	}
	var builder strings.Builder
	for y, row := range b.Matrix {
		for x, el := range row {
			if _, ok := has[Point{x, y}]; ok {
				builder.WriteString("•")
			} else {
				builder.WriteByte(el)
			}
		}
		builder.WriteByte('\n')
	}
	fmt.Println(builder.String())

}

type Graph map[Point]map[Point]int

type Path struct {
	Points []Point
	Cost   int
}

var ErrNoPath = errors.New("not path")

func (g Graph) path(start, end Point) (Path, error) {
	if len(g) == 0 {
		return Path{}, errors.New("cannot find path in empty graph")
	}
	if _, ok := g[start]; !ok {
		return Path{}, errors.New("cannot find start in graph")
	}
	if _, ok := g[end]; !ok {
		return Path{}, errors.New("cannot find end in graph")
	}

	explored := map[Point]bool{}
	previous := map[Point]Point{}

	queue := newQueue()

	path := Path{}
	var found bool
	var iterations int

	queue.set(start, 0)

	for !queue.isEmpty() {
		iterations++
		point, cost := queue.next()

		if point == end {
			path.Cost = cost
			for point != start {
				path.Points = append(path.Points, point)
				point = previous[point]
			}
			found = true
			break

		}

		explored[point] = true

		for nPoint, nCost := range g[point] {
			if explored[nPoint] {
				continue
			}

			if _, ok := queue.get(nPoint); !ok {
				previous[nPoint] = point
				queue.set(nPoint, cost+nCost)
				continue
			}

			frontierCost, _ := queue.get(nPoint)

			nodeCost := cost + nCost

			if nodeCost < frontierCost {
				previous[nPoint] = point
				queue.set(nPoint, nodeCost)
			}

		}
	}

	path.Points = append(path.Points, start)

	for i, j := 0, len(path.Points)-1; i < j; i, j = i+1, j-1 {
		path.Points[i], path.Points[j] = path.Points[j], path.Points[i]
	}

	if !found {
		return path, ErrNoPath
	}

	return path, nil
}

type Queue struct {
	points []Point
	nodes  map[Point]int
}

func (q *Queue) Len() int {
	return len(q.points)
}

func (q *Queue) Swap(i, j int) {
	q.points[i], q.points[j] = q.points[j], q.points[i]
}

func (q *Queue) Less(i, j int) bool {
	a := q.points[i]
	b := q.points[j]
	return q.nodes[a] < q.nodes[b]
}

func (q *Queue) set(point Point, priority int) {
	if _, ok := q.nodes[point]; !ok {
		q.points = append(q.points, point)
	}
	q.nodes[point] = priority
	sort.Sort(q)
}

func (q *Queue) next() (Point, int) {
	point, points := q.points[0], q.points[1:]
	q.points = points
	priority := q.nodes[point]
	delete(q.nodes, point)
	return point, priority
}

func (q *Queue) isEmpty() bool {
	return len(q.points) == 0
}

func (q *Queue) get(p Point) (int, bool) {
	priority, ok := q.nodes[p]
	return priority, ok
}

func newQueue() *Queue {
	return &Queue{
		nodes: map[Point]int{},
	}
}
