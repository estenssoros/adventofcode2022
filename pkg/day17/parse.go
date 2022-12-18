package day17

import (
	"context"
	"strconv"
)

func parseInput(ctx context.Context, input string) chan byte {
	ch := make(chan byte)
	go func() {
		defer close(ch)
		var i int
		for {
			if i > len(input)-1 {
				i = 0
			}
			next := input[i]
			i++
			select {
			case <-ctx.Done():
				return
			case ch <- next:
			}
		}
	}()
	return ch
}

var Rocks = []Rock{
	{
		{'0', '0', '1', '1', '1', '1', '0'},
	},
	{
		{'0', '0', '0', '1', '0', '0', '0'},
		{'0', '0', '1', '1', '1', '0', '0'},
		{'0', '0', '0', '1', '0', '0', '0'},
	},
	{
		{'0', '0', '0', '0', '1', '0', '0'},
		{'0', '0', '0', '0', '1', '0', '0'},
		{'0', '0', '1', '1', '1', '0', '0'},
	},
	{
		{'0', '0', '1', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0', '0', '0'},
	},
	{
		{'0', '0', '1', '1', '0', '0', '0'},
		{'0', '0', '1', '1', '0', '0', '0'},
	},
}

func BinaryToInt(row []byte) int {
	out, err := strconv.ParseInt(string(row), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(out)
}
