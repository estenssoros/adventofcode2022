package day11

import (
	"strconv"
)

type Monkey struct {
	ID             int
	Items          []int64 // list of worry level
	Operation      Operation
	Test           Test
	ItemsInspected int
}

type Test struct {
	Divisible int64
	TrueDst   int
	FalseDst  int
}

type Operation struct {
	A  string
	B  string
	Op string
}

func (o Operation) do(item int64) int64 {
	var a, b int64 = item, 0
	if o.B == "old" {
		b = item
	} else {
		b = mustParseInt(o.B)
	}
	switch o.Op {
	case "*":
		return a * b
	case "-":
		return a - b
	case "+":
		return a + b
	default:
		panic(o.Op)
	}
}

func (m *Monkey) InspectItems(worry func(int64) int64) []int64 {
	out := make([]int64, len(m.Items))
	for i, item := range m.Items {
		out[i] = worry(m.Operation.do(item))
		m.ItemsInspected++
	}
	m.Items = nil
	return out
}

func mustParseInt(s string) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int64(i)
}

func (m Monkey) WhereThrow(i int64) int {
	if i%m.Test.Divisible == 0 {
		return m.Test.TrueDst
	}
	return m.Test.FalseDst
}

func (m *Monkey) Inspect(monkeys []*Monkey, worry func(int64) int64) {
	var d = 1
	for _, mm := range monkeys {
		d *= int(mm.Test.Divisible)
	}
	for _, item := range m.Items {
		level := m.Operation.do(item)
		level = worry(level)
		level = level % int64(d)
		dest := m.WhereThrow(level)
		monkeys[dest].Items = append(monkeys[dest].Items, level)
	}
	m.ItemsInspected += len(m.Items)
	m.Items = []int64{}
}
