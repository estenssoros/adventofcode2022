package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreRounds(t *testing.T) {
	rounds := []Round{
		{
			Player:   'Y',
			Opponent: 'A',
		},
		{
			Player:   'X',
			Opponent: 'B',
		},
		{
			Player:   'Z',
			Opponent: 'C',
		},
	}
	assert.Equal(t, 15, scoreRoundsPart1(rounds))
}

func TestScoreRound(t *testing.T) {
	{
		round := Round{
			Opponent: 'A',
			Player:   'Y',
		}
		assert.Equal(t, 8, round.ScorePart1())
	}
	{
		round := Round{
			Opponent: 'B',
			Player:   'X',
		}
		assert.Equal(t, 1, round.ScorePart1())
	}
	{
		round := Round{
			Opponent: 'C',
			Player:   'Z',
		}
		assert.Equal(t, 6, round.ScorePart1())
	}
}

func TestRunes(t *testing.T) {
	assert.Equal(t, 'X', 'A'+23)
}

func TestPlayerEqualOpponent(t *testing.T) {
	assert.True(t, playerEqualOpponent(PlayerRock, OpponentRock))
	assert.True(t, playerEqualOpponent(PlayerPaper, OpponentPaper))
	assert.True(t, playerEqualOpponent(PlayerScissors, OpponentScissors))
}
