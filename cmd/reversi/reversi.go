package reversi

import (
	reversi_errors "develop/cmd/error"
	"fmt"
)

const (
	Empty = iota
	Black
	White
)

const (
	upperLine               = "upperLine"
	rightDiagonalUpperLine  = "rightDiagonalUpperLine"
	rightLine               = "rightLine"
	rightDiagonalBottomLine = "rightDiagonalBottomLine"
	bottomLine              = "bottomLine"
	leftDiagonalBottomLine  = "leftDiagonalBottomLine"
	leftLine                = "leftLine"
	leftDiagonalUpperLine   = "leftDiagonalUpperLine"
)

type Reversi struct {
	Board        [][]int
	StoneCount   int
	CurrentStone int
}

func InitializeReversi() Reversi {
	var board [][]int
	for {
		board = append(board, []int{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty})
		if len(board) == 8 {
			break
		}
	}

	board[3][3] = Black
	board[3][4] = White
	board[4][3] = White
	board[4][4] = Black

	return Reversi{board, 60, Black}
}

func (rev *Reversi) CountStone() (black, white int) {
	for _, line := range rev.Board {
		for _, cell := range line {
			if cell == 1 {
				black++
			}
			if cell == 2 {
				white++
			}
		}
	}
	return
}

// 0の場合はDraw
func (rev *Reversi) Judgement() int {
	black, white := rev.CountStone()
	if black > white {
		return Black
	}

	if black < white {
		return White
	}

	return 0
}

func (rev *Reversi) DisplayBoard() {
	fmt.Println(60 - rev.StoneCount)
	fmt.Println("====================================")
	for _, line := range rev.Board {
		fmt.Println(line)
	}
	fmt.Println("====================================")
}

type ReversiPlayer struct {
	MyStone       int
	OpponentStone int
	Rev           *Reversi
}

func (rp *ReversiPlayer) judgeStone(x, y int) (map[string]int, bool) {
	if rp.Rev.Board[y][x] != Empty {
		return nil, false
	}

	lineCountMap := make(map[string]int)
	lineCountMap[upperLine] = 0
	lineCountMap[rightDiagonalUpperLine] = 0
	lineCountMap[rightLine] = 0
	lineCountMap[rightDiagonalBottomLine] = 0
	lineCountMap[bottomLine] = 0
	lineCountMap[leftDiagonalBottomLine] = 0
	lineCountMap[leftLine] = 0
	lineCountMap[leftDiagonalUpperLine] = 0

	// 上側に石を置けるか確認
	if y != 0 && rp.Rev.Board[y-1][x] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y-i][x] == rp.OpponentStone {
			lineCountMap[upperLine]++
			if y-i == 0 {
				lineCountMap[upperLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y-i][x] == Empty {
			lineCountMap[upperLine] = 0
		}
	}

	// 右上側に石を置けるか確認
	if x != 7 && y != 0 && rp.Rev.Board[y-1][x+1] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y-i][x+i] == rp.OpponentStone {
			lineCountMap[rightDiagonalUpperLine]++
			if x+i == 7 || y-i == 0 {
				lineCountMap[rightDiagonalUpperLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y-i][x+i] == Empty {
			lineCountMap[rightDiagonalUpperLine] = 0
		}
	}

	// 右側に石を置けるか確認
	if x != 7 && rp.Rev.Board[y][x+1] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y][x+i] == rp.OpponentStone {
			lineCountMap[rightLine]++
			if x+i == 7 {
				lineCountMap[rightLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y][x+i] == Empty {
			lineCountMap[rightLine] = 0
		}
	}

	// 右下側に石を置けるか確認
	if x != 7 && y != 7 && rp.Rev.Board[y+1][x+1] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y+i][x+i] == rp.OpponentStone {
			lineCountMap[rightDiagonalBottomLine]++
			if x+i == 7 || y+i == 7 {
				lineCountMap[rightDiagonalBottomLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y+i][x+i] == Empty {
			lineCountMap[rightDiagonalBottomLine] = 0
		}
	}

	// 下側に石を置けるか確認
	if y != 7 && rp.Rev.Board[y+1][x] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y+i][x] == rp.OpponentStone {
			lineCountMap[bottomLine]++
			if y+i == 7 {
				lineCountMap[bottomLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y+i][x] == Empty {
			lineCountMap[bottomLine] = 0
		}
	}

	// 左下側に石を置けるか確認
	if x != 0 && y != 7 && rp.Rev.Board[y+1][x-1] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y+i][x-i] == rp.OpponentStone {
			lineCountMap[leftDiagonalBottomLine]++
			if x-i == 0 || y+i == 7 {
				lineCountMap[leftDiagonalBottomLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y+i][x-i] == Empty {
			lineCountMap[leftDiagonalBottomLine] = 0
		}
	}

	// 左側に石を置けるか確認
	if x != 0 && rp.Rev.Board[y][x-1] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y][x-i] == rp.OpponentStone {
			lineCountMap[leftLine]++
			if x-i == 0 {
				lineCountMap[leftLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y][x-i] == Empty {
			lineCountMap[leftLine] = 0
		}
	}

	// 左上側に石を置けるか確認
	if x != 0 && y != 0 && rp.Rev.Board[y-1][x-1] == rp.OpponentStone {
		i := 1
		for rp.Rev.Board[y-i][x-i] == rp.OpponentStone {
			lineCountMap[leftDiagonalUpperLine]++
			if x-1 == 0 || y-i == 0 {
				lineCountMap[leftDiagonalUpperLine] = 0
				break
			}
			i++
		}

		if rp.Rev.Board[y-i][x-i] == Empty {
			lineCountMap[leftDiagonalUpperLine] = 0
		}
	}

	for _, count := range lineCountMap {
		if count != 0 {
			return lineCountMap, true
		}
	}

	return nil, false
}

func (rp *ReversiPlayer) reverseStone(x, y int, lineCountMap map[string]int) {
	for line, count := range lineCountMap {
		if count == 0 {
			continue
		}

		switch line {
		case upperLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y-i][x] = rp.MyStone
			}
		case rightDiagonalUpperLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y-i][x+i] = rp.MyStone
			}
		case rightLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y][x+i] = rp.MyStone
			}
		case rightDiagonalBottomLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y+i][x+i] = rp.MyStone
			}
		case bottomLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y+i][x] = rp.MyStone
			}
		case leftDiagonalBottomLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y+i][x-i] = rp.MyStone
			}
		case leftLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y][x-i] = rp.MyStone
			}
		case leftDiagonalUpperLine:
			for i := 1; i <= count; i++ {
				rp.Rev.Board[y-i][x-i] = rp.MyStone
			}
		}
	}
}

func (rp *ReversiPlayer) SetStone(x, y int) error {
	if rp.Rev.CurrentStone != rp.MyStone {
		return &reversi_errors.SetStoneError{
			Player: rp.MyStone,
			Detail: "Not my turn.",
		}
	}

	lineCountMap, setFlag := rp.judgeStone(x, y)
	if !setFlag {
		return &reversi_errors.SetStoneError{
			Player: rp.MyStone,
			Detail: fmt.Sprintf("You can't put it in this cell. x: %v y: %v", x, y),
		}
	}

	rp.Rev.Board[y][x] = rp.MyStone
	rp.reverseStone(x, y, lineCountMap)
	rp.Rev.CurrentStone = rp.OpponentStone
	rp.Rev.StoneCount--

	return nil
}
