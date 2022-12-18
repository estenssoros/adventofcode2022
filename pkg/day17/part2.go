package day17

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part2(input, 1000000000000)
		if err != nil {
			return errors.Wrap(err, "part2")
		}
		fmt.Println("part2:", out)
		return nil
	},
}

var mapping = map[int]byte{
	1:   'q',
	2:   'K',
	3:   'r',
	4:   'f',
	6:   'g',
	7:   'L',
	8:   'b',
	11:  'z',
	14:  'D',
	15:  'p',
	16:  'i',
	17:  'w',
	18:  'v',
	20:  'e',
	23:  'u',
	28:  'c',
	30:  'a',
	31:  'A',
	36:  'J',
	40:  'G',
	42:  'F',
	48:  'y',
	49:  'x',
	52:  's',
	56:  'j',
	60:  'h',
	61:  'I',
	62:  't',
	63:  'E',
	68:  'n',
	70:  'm',
	79:  'o',
	80:  'C',
	102: 'l',
	119: 'H',
	120: 'B',
	124: 'd',
	126: 'k',
}

func part2(input string, numIterations int) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rockChan := RockChan(ctx, Rocks)
	cave := Cave{
		wind: parseInput(ctx, input),
	}
	bar := progressbar.Default(int64(2022))
	for i := 0; i < 2022; i++ {
		bar.Add(1)
		rock := <-rockChan
		cave.grow(3)
		cave.dropRock(rock)
	}
	bar.Close()
	builder := strings.Builder{}
	for _, row := range cave.grid {
		b, ok := mapping[BinaryToInt(row)]
		if !ok {
			panic(fmt.Sprintf("missing %d", BinaryToInt(row)))
		}
		builder.WriteByte(b)
	}
	caveString := builder.String()
	match := findPattern(builder.String())
	fmt.Printf("found repeating pattern: of %d\n", len(match))
	fmt.Println(match)
	for !strings.HasSuffix(caveString, match) {
		if len(caveString) == 0 {
			return 0, errors.New("could not match cave string")
		}
		caveString = caveString[:len(caveString)-1]
	}
	caveLen, matchLen := len(caveString), len(match)
	for caveLen+matchLen < numIterations {
		caveLen += matchLen
	}
	fmt.Println(caveLen)
	// for !strings.HasSuffix(data,)
	return len(cave.grid), nil
}
