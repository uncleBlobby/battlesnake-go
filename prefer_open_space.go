package main

// look at each possible move direction
// look from there to see how far you can see :)

func DetermineOpenSpaces(state GameState, lookDistance map[string]int) map[string]int {
	myHead := state.You.Body[0]
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	for i := myHead.X - 1; i > 0; i-- {
		var targetCoord = Coord{myHead.X - 1 - lookDistance["left"], myHead.Y}
		if !checkAnyGivenCoordForSnakeBody(state, targetCoord) {
			lookDistance["left"]++
		} else {
			break
		}
	}
	for i := myHead.X + 1; i < boardWidth-1; i++ {
		var targetCoord = Coord{myHead.X + 1 + lookDistance["right"], myHead.Y}
		if !checkAnyGivenCoordForSnakeBody(state, targetCoord) {
			lookDistance["right"]++
		} else {
			break
		}
	}
	for i := myHead.Y + 1; i < boardHeight-1; i++ {
		var targetCoord = Coord{myHead.X, myHead.Y + 1 + lookDistance["up"]}
		if !checkAnyGivenCoordForSnakeBody(state, targetCoord) {
			lookDistance["up"]++
		} else {
			break
		}
	}
	for i := myHead.Y - 1; i > 0; i-- {
		var targetCoord = Coord{myHead.X, myHead.Y - 1 - lookDistance["down"]}
		if !checkAnyGivenCoordForSnakeBody(state, targetCoord) {
			lookDistance["down"]++
		} else {
			break
		}
	}

	return lookDistance
}

func checkAnyGivenCoordForSnakeBody(state GameState, target Coord) bool {
	snakes := state.Board.Snakes
	for i := 0; i < len(snakes); i++ {
		for j := 0; j < int(snakes[i].Length); j++ {
			if (target.X == snakes[i].Body[j].X) && (target.Y == snakes[i].Body[j].Y) {
				return true
			}
		}
	}
	return false
}

func ResetLookDistance(lookDistance map[string]int) map[string]int {
	lookDistance["left"] = 0
	lookDistance["right"] = 0
	lookDistance["up"] = 0
	lookDistance["down"] = 0

	return lookDistance
}
