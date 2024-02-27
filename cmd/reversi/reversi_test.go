package reversi_test

import (
	"develop/cmd/reversi"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	player = reversi.ReversiPlayer{
		MyStone:       reversi.Black,
		OpponentStone: reversi.White,
	}

	computer = reversi.ReversiPlayer{
		MyStone:       reversi.White,
		OpponentStone: reversi.Black,
	}
)

func TestSetStone_NormalcyCheck(t *testing.T) {
	var board [][]int
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 2, 1, 1, 2, 1})
	board = append(board, []int{1, 2, 1, 2, 1, 2, 1, 1})
	board = append(board, []int{1, 1, 2, 2, 2, 1, 1, 1})
	board = append(board, []int{1, 2, 2, 0, 2, 2, 2, 1})
	board = append(board, []int{1, 1, 2, 2, 2, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})

	rev := reversi.Reversi{
		Board:        board,
		StoneCount:   30,
		CurrentStone: 1,
	}

	player.Rev = &rev
	computer.Rev = &rev

	err := player.SetStone(3, 5)
	assert.Nil(t, err)

	assert.Equal(t, rev.Board[3][2], 1)
	assert.Equal(t, rev.Board[6][2], 1)
	assert.Equal(t, rev.Board[1][3], 1)
	assert.Equal(t, rev.Board[3][3], 1)
	assert.Equal(t, rev.Board[5][3], 1)
	assert.Equal(t, rev.Board[2][4], 1)
	assert.Equal(t, rev.Board[3][4], 1)
	assert.Equal(t, rev.Board[4][4], 1)
	assert.Equal(t, rev.Board[1][5], 1)
	assert.Equal(t, rev.Board[2][5], 1)
	assert.Equal(t, rev.Board[4][5], 1)
	assert.Equal(t, rev.Board[5][5], 1)
	assert.Equal(t, rev.Board[6][5], 1)
	assert.Equal(t, rev.Board[2][6], 1)
	assert.Equal(t, rev.Board[3][6], 1)
	assert.Equal(t, rev.Board[4][6], 1)
}

func TestCountStone(t *testing.T) {
	var board [][]int
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 2, 1, 1, 2, 1})
	board = append(board, []int{1, 2, 1, 2, 1, 2, 1, 1})
	board = append(board, []int{1, 1, 2, 2, 2, 1, 1, 1})
	board = append(board, []int{1, 2, 2, 1, 2, 2, 2, 1})
	board = append(board, []int{1, 1, 2, 2, 2, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})

	rev := reversi.Reversi{
		Board:        board,
		StoneCount:   30,
		CurrentStone: 1,
	}

	black, white := rev.CountStone()
	assert.Equal(t, black, 48)
	assert.Equal(t, white, 16)
}

func TestJudgement(t *testing.T) {
	var board [][]int
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 2, 1, 1, 2, 1})
	board = append(board, []int{1, 2, 1, 2, 1, 2, 1, 1})
	board = append(board, []int{1, 1, 2, 2, 2, 1, 1, 1})
	board = append(board, []int{1, 2, 2, 1, 2, 2, 2, 1})
	board = append(board, []int{1, 1, 2, 2, 2, 1, 1, 1})
	board = append(board, []int{1, 1, 1, 1, 1, 1, 1, 1})

	rev := reversi.Reversi{
		Board:        board,
		StoneCount:   30,
		CurrentStone: 1,
	}

	result := rev.Judgement()
	assert.Equal(t, result, reversi.Black)
}
