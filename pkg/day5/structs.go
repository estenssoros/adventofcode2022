package day5

import (
	"fmt"
	"strings"
)

type Stack struct {
	boxes []byte
}

func (s Stack) String() string {
	b := strings.Builder{}
	for _, box := range s.boxes {
		b.WriteByte(box)
	}
	return b.String()
}

func (s *Stack) push(box byte) {
	s.boxes = append(s.boxes, box)
}

func (s *Stack) pop() byte {
	box := s.boxes[len(s.boxes)-1]
	s.boxes = s.boxes[:len(s.boxes)-1]
	return box
}

func (s *Stack) peek() byte {
	return s.boxes[len(s.boxes)-1]
}

type Ship struct {
	Containers []*Stack
	Operations []*Operation
}

type Operation struct {
	Num int
	Src int
	Dst int
}

func (o Operation) String() string {
	return fmt.Sprintf("move %d from %d to %d", o.Num, o.Src, o.Dst)
}
