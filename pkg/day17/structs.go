package day17

import (
	"context"
	"fmt"
	"strings"
)

var (
	Empty   byte = '0'
	Falling byte = '@'
	Filled  byte = '1'
	Right   byte = '>'
	Left    byte = '<'
)

type Rock [][]byte

func (r Rock) String() string {
	b := strings.Builder{}
	for _, row := range r {
		b.WriteString(string(row) + "\n")
	}
	return strings.TrimSpace(b.String())
}

func ShiftRight(r [][]byte) {
	// fmt.Println("blowing right")
	if colEmpty(r, 6) {
		for y := 0; y < len(r); y++ {
			for x := 6; x > 0; x-- {
				r[y][x] = r[y][x-1]
			}
			r[y][0] = Empty
		}
	}
}

func ShiftLeft(r [][]byte) {
	// fmt.Println("blowing left")
	if colEmpty(r, 0) {
		for y := 0; y < len(r); y++ {
			for x := 0; x < 6; x++ {
				r[y][x] = r[y][x+1]
			}
			r[y][6] = Empty
		}
	}
}

func colEmpty(r [][]byte, col int) bool {
	for row := range r {
		if r[row][col] != Empty {
			return false
		}
	}
	return true
}
func copyRock(r [][]byte) [][]byte {
	out := make([][]byte, len(r))
	for i, row := range r {
		out[i] = copyRow(row)
	}
	return out
}
func copyRow(row []byte) []byte {
	out := make([]byte, len(row))
	for i, el := range row {
		out[i] = el
	}
	return out
}

func RockChan(ctx context.Context, rocks []Rock) <-chan Rock {
	ch := make(chan Rock)
	rockLen := len(rocks)
	var i int
	go func() {
		defer close(ch)
		for {
			if i >= rockLen {
				i = 0
			}
			r := rocks[i]
			select {
			case <-ctx.Done():
				return
			case ch <- copyRock(r):
				i++
			}
		}
	}()
	return ch
}

type Cave struct {
	grid [][]byte
	wind <-chan byte
}

func (c *Cave) grow(num int) {
	for i := 0; i < num; i++ {
		c.grid = append(c.grid, emptyRow(7))
	}
}

func (c Cave) String() string {
	b := strings.Builder{}
	for i := len(c.grid) - 1; i >= 0; i-- {
		b.WriteString(fmt.Sprint(i) + "\t" + string(c.grid[i]) + "\t" + fmt.Sprint(BinaryToInt(c.grid[i])) + "\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func emptyRow(width int) []byte {
	row := make([]byte, width)
	for i := 0; i < width; i++ {
		row[i] = Empty
	}
	return row
}

func (c *Cave) dropRock(rock Rock) {
	c.grow(len(rock))
	rockBottom := c.findRockBottom(rock)
	for y, row := range rock {
		for x := 0; x < len(row); x++ {
			if c.grid[rockBottom+len(rock)-1-y][x] == Empty {
				c.grid[rockBottom+len(rock)-1-y][x] = row[x]
			}
		}
	}
	c.trimGrid()
}

func (c *Cave) findRockBottom(rock [][]byte) int {
	y := len(c.grid) - len(rock)
	// fmt.Printf("new rock @ %d!\n", y)
	// fmt.Println(Rock(rock))
	c.Blow(rock, y)
	for y > 0 {
		// fmt.Println(Rock(rock))
		if !c.canShiftDown(rock, y) {
			return y
		}
		y--
		c.Blow(rock, y)
	}
	return 0
}

func (c *Cave) Blow(rock [][]byte, y int) {
	dir := <-c.wind
	switch dir {
	case Right:
		if c.canShiftRight(rock, y) {
			ShiftRight(rock)
		}
	case Left:
		if c.canShiftLeft(rock, y) {
			ShiftLeft(rock)
		}
	default:
		panic("unknown wind: " + string(dir))
	}
}

func (c *Cave) canShiftDown(rock [][]byte, y int) bool {
	for i := len(rock) - 1; i >= 0; i-- {
		for x, r := range rock[i] {
			if r == Filled && c.grid[y+(len(rock)-i-1)-1][x] == Filled {
				return false
			}
		}
	}
	return true
}

func (c *Cave) canShiftRight(rock [][]byte, row int) bool {
	for y := len(rock) - 1; y >= 0; y-- {
		for x := 0; x < 7-1; x++ {
			r := rock[y][x]
			if r == Filled && c.grid[row+(len(rock)-y-1)][x+1] == Filled {
				return false
			}
		}
	}
	return true
}

func (c *Cave) canShiftLeft(rock [][]byte, row int) bool {
	for y := len(rock) - 1; y >= 0; y-- {
		for x := 7 - 1; x > 0; x-- {
			r := rock[y][x]
			if r == Filled && c.grid[row+(len(rock)-y-1)][x-1] == Filled {
				return false
			}
		}
	}
	return true
}

func (c *Cave) trimGrid() {
	grid := [][]byte{}
	for _, row := range c.grid {
		if isEmpty(row) {
			continue
		}
		grid = append(grid, row)
	}
	c.grid = grid
}

func isEmpty(row []byte) bool {
	for _, el := range row {
		if el != Empty {
			return false
		}
	}
	return true
}
