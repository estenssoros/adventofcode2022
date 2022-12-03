package day2

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	RoundLose = 'X'
	RoundDraw = 'Y'
	RoundWin  = 'Z'
)

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part2() },
}

func part2() error {
	rounds, err := parseInput(input)
	if err != nil {
		return errors.Wrap(err, "parseInput")
	}
	fmt.Println(scoreRoundsPart2(rounds))
	return nil
}

func scoreRoundsPart2(rounds []Round) int {
	var score int
	for _, round := range rounds {
		score += round.ScorePart2()
	}
	return score
}

func (r *Round) ScorePart2() int {
	switch r.Player {
	case RoundDraw:
		r.Player = opponentToPlayer(r.Opponent)
	case RoundWin:
		if r.Opponent == OpponentPaper {
			r.Player = PlayerScissors
		} else if r.Opponent == OpponentRock {
			r.Player = PlayerPaper
		} else if r.Opponent == OpponentScissors {
			r.Player = PlayerRock
		}
	case RoundLose:
		if r.Opponent == OpponentPaper {
			r.Player = PlayerRock
		} else if r.Opponent == OpponentRock {
			r.Player = PlayerScissors
		} else if r.Opponent == OpponentScissors {
			r.Player = PlayerPaper
		}
	}
	return r.ScorePart1()
}

func (r Round) OutcomePart2() int {
	return r.ScorePart2()
}
