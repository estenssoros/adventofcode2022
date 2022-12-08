package day8

type Matrix struct {
	Cells  [][]int
	Height int
	Width  int
}

func (m Matrix) valid(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x > m.Width-1 || y > m.Height-1 {
		return false
	}
	return true
}

func (m Matrix) get(x, y int) int {
	return m.Cells[y][x]
}
