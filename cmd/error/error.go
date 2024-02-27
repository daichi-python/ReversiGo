package reversi_errors

import "fmt"

type SetStoneError struct {
	Player int
	Detail string
}

func (sse *SetStoneError) Error() string {
	var player string
	switch sse.Player {
	case 1:
		player = "black"
	case 2:
		player = "white"
	}

	return fmt.Sprintf("Error occurred when playing %s stone player. Detail: %s", player, sse.Detail)
}
