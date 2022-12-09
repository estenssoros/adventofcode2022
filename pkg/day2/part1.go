package day2

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Round struct {
	Opponent rune
	Player   rune
}

const (
	OpponentRock     = 'A'
	OpponentPaper    = 'B'
	OpponentScissors = 'C'
	PlayerRock       = 'X'
	PlayerPaper      = 'Y'
	PlayerScissors   = 'Z'

	RoundDrawScore = 3
	RoundWinScore  = 6
	RoundLoseScore = 0
)

func playerEqualOpponent(p, o rune) bool {
	return o+23 == p
}

func playerToOpponent(p rune) rune {
	return p - 23
}

func opponentToPlayer(o rune) rune {
	return o + 23
}

func (r Round) ScorePart1() int {
	var score int
	switch r.Player {
	case PlayerRock:
		score += 1
	case PlayerPaper:
		score += 2
	case PlayerScissors:
		score += 3
	}
	return score + r.Outcome()
}

func (r Round) Outcome() int {
	if playerEqualOpponent(r.Player, r.Opponent) {
		return RoundDrawScore
	}

	if r.Opponent == OpponentRock {
		if r.Player == PlayerPaper {
			return RoundWinScore
		} else {
			return RoundLoseScore
		}
	}

	if r.Opponent == OpponentPaper {
		if r.Player == PlayerRock {
			return RoundLoseScore
		} else {
			return RoundWinScore
		}
	}

	if r.Opponent == OpponentScissors {
		if r.Player == PlayerPaper {
			return RoundLoseScore
		} else {
			return RoundWinScore
		}
	}
	panic("asdf")
	return -1
}

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := part1(input)
		if err != nil {
			return errors.Wrap(err, "part1")
		}
		fmt.Println(out)
		return nil
	},
}

func part1(input string) (int, error) {
	rounds, err := parseInput(input)
	if err != nil {
		return 0, errors.Wrap(err, "parseInput")
	}

	return scoreRoundsPart1(rounds), nil
}

func scoreRoundsPart1(rounds []Round) int {
	var score int
	for _, round := range rounds {
		score += round.ScorePart1()
	}
	return score
}

func parseInput(input string) ([]Round, error) {
	rounds := []Round{}
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, errors.Errorf("could not parse: %s", line)
		}
		rounds = append(rounds, Round{rune(fields[0][0]), rune(fields[1][0])})
	}

	return rounds, nil
}
