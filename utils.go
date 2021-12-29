package main

func FindNextCoord(direction string, startPos Coord) Coord {
	var pos Coord
	switch direction {
	case "lower-left":
		pos.X = startPos.X - 1
		pos.Y = startPos.Y - 1
	case "left":
		pos.X = startPos.X - 1
		pos.Y = startPos.Y
	case "upper-left":
		pos.X = startPos.X - 1
		pos.Y = startPos.Y + 1
	case "up":
		pos.X = startPos.X
		pos.Y = startPos.Y + 1
	case "upper-right":
		pos.X = startPos.X + 1
		pos.Y = startPos.Y + 1
	case "right":
		pos.X = startPos.X + 1
		pos.Y = startPos.Y
	case "lower-right":
		pos.X = startPos.X + 1
		pos.Y = startPos.Y - 1
	case "down":
		pos.X = startPos.X
		pos.Y = startPos.Y - 1
	}

	return pos
}

func FindEnemySnakes(state GameState) []Battlesnake {
	me := state.You
	//fmt.Println("Entered Find Enemy Snakes funciton")
	enemySnakes := []Battlesnake{}
	//fmt.Println("Number of snakes on board: ", len(state.Board.Snakes))
	for i := 0; i < len(state.Board.Snakes); i++ {
		//fmt.Println("Found snakes: ", state.Board.Snakes)
		if state.Board.Snakes[i].ID != me.ID {
			enemySnakes = append(enemySnakes, state.Board.Snakes[i])
		}
	}
	//fmt.Println("enemySnakes inside Find function: ", enemySnakes)
	return enemySnakes
}

func CheckIfLargerSnakesNearMyHead(enemies []Battlesnake,
	state GameState,
	nearMyHead map[string]bool) map[string]bool {
	myHead := state.You.Body[0]
	myLength := state.You.Length

	for i := 0; i < len(enemies); i++ {
		if (FindNextCoord("left", FindNextCoord("left", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("up", FindNextCoord("left", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("down", FindNextCoord("left", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) {
			nearMyHead["left"] = true
		}
		if (FindNextCoord("left", FindNextCoord("up", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("up", FindNextCoord("up", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("right", FindNextCoord("up", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) {
			nearMyHead["up"] = true
		}
		if (FindNextCoord("down", FindNextCoord("right", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("up", FindNextCoord("right", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("right", FindNextCoord("right", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) {
			nearMyHead["right"] = true
		}
		if (FindNextCoord("down", FindNextCoord("down", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("left", FindNextCoord("down", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) ||
			(FindNextCoord("right", FindNextCoord("down", myHead)) == enemies[i].Body[0] && myLength <= enemies[i].Length) {
			nearMyHead["down"] = true
		}
	}

	return nearMyHead
}

func CheckIfSmallerSnakesNearMyHead(enemies []Battlesnake,
	state GameState,
	nearMyHead map[string]bool) map[string]bool {
	myHead := state.You.Body[0]
	myLength := state.You.Length

	for i := 0; i < len(enemies); i++ {
		if (FindNextCoord("left", FindNextCoord("left", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("up", FindNextCoord("left", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("down", FindNextCoord("left", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) {
			nearMyHead["left"] = true
		}
		if (FindNextCoord("left", FindNextCoord("up", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("up", FindNextCoord("up", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("right", FindNextCoord("up", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) {
			nearMyHead["up"] = true
		}
		if (FindNextCoord("down", FindNextCoord("right", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("up", FindNextCoord("right", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("right", FindNextCoord("right", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) {
			nearMyHead["right"] = true
		}
		if (FindNextCoord("down", FindNextCoord("down", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("left", FindNextCoord("down", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) ||
			(FindNextCoord("right", FindNextCoord("down", myHead)) == enemies[i].Body[0] && myLength > enemies[i].Length) {
			nearMyHead["down"] = true
		}
	}

	return nearMyHead
}
