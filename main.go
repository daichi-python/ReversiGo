package main

import "develop/cmd/reversi"

func main() {
	rev := reversi.InitializeReversi()
	player := reversi.ReversiPlayer{
		MyStone:       reversi.Black,
		OpponentStone: reversi.White,
		Rev:           &rev,
	}

	computer := reversi.ReversiPlayer{
		MyStone:       reversi.White,
		OpponentStone: reversi.Black,
		Rev:           &rev,
	}

	player.SetStone(3, 5)
	computer.SetStone(2, 5)
	player.SetStone(5, 3)
	computer.SetStone(3, 6)
	player.SetStone(3, 7)
	computer.SetStone(5, 2)
	player.SetStone(5, 1)
	computer.SetStone(6, 3)
	player.SetStone(7, 3)
	computer.SetStone(6, 1)
	player.SetStone(7, 1)
	computer.SetStone(4, 7)
}
