package day10

import (
	"strings"
)

type Operation struct {
	IsNoop bool
	AddX   int
}

type CRT struct {
	Screen         [][]byte
	SpritePosition int
	width          int
}

func newCRT(width, height int) CRT {
	return CRT{
		SpritePosition: 0,
		width:          width,
	}
}

func (c CRT) String() string {
	b := strings.Builder{}
	for _, row := range c.Screen {
		for _, cell := range row {
			b.WriteByte(cell)
		}
		b.WriteByte('\n')
	}
	return strings.TrimSpace(b.String())
}

func (c *CRT) runOperations(operations []Operation) {
	var x, cycle = 1, 1
	row := []byte{}
	tick := func() {
		if (cycle-1)%c.width == 0 {
			c.Screen = append(c.Screen, row)
			row = nil
			c.SpritePosition = 0
		}
		if c.SpritePosition >= x-1 && c.SpritePosition <= x+1 {
			row = append(row, '#')
		} else {
			row = append(row, '.')
		}
		cycle++
		c.SpritePosition++
	}
	for _, operation := range operations {
		if operation.IsNoop {
			tick()
		} else {
			tick()
			tick()
			x += operation.AddX
		}
	}
	c.Screen = append(c.Screen, row)

}
