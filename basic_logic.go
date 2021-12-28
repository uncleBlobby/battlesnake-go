package main

import "math"

func findFoodDistances(food []Coord, myHead Coord) Food {

	foodList := []Food{}
	for i := 0; i < len(food); i++ {
		var newFood Food
		newFood.X = food[i].X
		newFood.Y = food[i].Y
		newFood.Distance.X = int(math.Abs(float64(food[i].X - myHead.X)))
		newFood.Distance.Y = int(math.Abs(float64(food[i].Y - myHead.Y)))
		newFood.Distance.Total = newFood.Distance.X + newFood.Distance.Y
		foodList = append(foodList, newFood)
	}

	// sort food list so first member is closest
	closestFood := foodList[0]
	for i := 1; i < len(foodList); i++ {
		if foodList[i].Distance.Total < closestFood.Distance.Total {
			closestFood = foodList[i]
		}
	}

	// can also return foodList (foodlist with distances)
	// removed for the time being

	return closestFood
}

func AvoidMyNeck(state GameState, possibleMoves map[string]bool) map[string]bool {
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")
	if myNeck.X < myHead.X {
		possibleMoves["left"] = false
	} else if myNeck.X > myHead.X {
		possibleMoves["right"] = false
	} else if myNeck.Y < myHead.Y {
		possibleMoves["down"] = false
	} else if myNeck.Y > myHead.Y {
		possibleMoves["up"] = false
	}

	return possibleMoves
}

func AvoidWalls(state GameState, possibleMoves map[string]bool) map[string]bool {
	myHead := state.You.Body[0]
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	if myHead.X == 0 {
		possibleMoves["left"] = false
	}
	if myHead.X == boardWidth-1 {
		possibleMoves["right"] = false
	}
	if myHead.Y == 0 {
		possibleMoves["down"] = false
	}
	if myHead.Y == boardHeight-1 {
		possibleMoves["up"] = false
	}

	return possibleMoves
}

func AvoidSelf(state GameState, possibleMoves map[string]bool) map[string]bool {
	myHead := state.You.Body[0]
	myBody := state.You.Body

	for i := 0; i < len(myBody); i++ {
		if myHead.X-1 == myBody[i].X && myHead.Y == myBody[i].Y {
			possibleMoves["left"] = false
		}
		if myHead.X+1 == myBody[i].X && myHead.Y == myBody[i].Y {
			possibleMoves["right"] = false
		}
		if myHead.X == myBody[i].X && myHead.Y-1 == myBody[i].Y {
			possibleMoves["down"] = false
		}
		if myHead.X == myBody[i].X && myHead.Y+1 == myBody[i].Y {
			possibleMoves["up"] = false
		}
	}

	return possibleMoves
}

func AvoidOtherSnakes(state GameState, possibleMoves map[string]bool) map[string]bool {
	snakes := state.Board.Snakes
	myHead := state.You.Body[0]

	for i := 0; i < len(snakes); i++ {
		for j := 0; j < int(snakes[i].Length); j++ {
			if (myHead.X-1 == snakes[i].Body[j].X) && (myHead.Y == snakes[i].Body[j].Y) {
				possibleMoves["left"] = false
			}
			if (myHead.X+1 == snakes[i].Body[j].X) && (myHead.Y == snakes[i].Body[j].Y) {
				possibleMoves["right"] = false
			}
			if (myHead.X == snakes[i].Body[j].X) && (myHead.Y-1 == snakes[i].Body[j].Y) {
				possibleMoves["down"] = false
			}
			if (myHead.X == snakes[i].Body[j].X) && (myHead.Y+1 == snakes[i].Body[j].Y) {
				possibleMoves["up"] = false
			}
		}
	}

	return possibleMoves
}

func PreferFoodMoves(state GameState, preferredMoves map[string]bool) map[string]bool {
	myHead := state.You.Body[0]

	if len(state.Board.Food) > 0 {
		closestFood := findFoodDistances(state.Board.Food, myHead)
		//fmt.Println("Closest Food:", closestFood)
		if closestFood.X < myHead.X {
			preferredMoves["left"] = true
		}
		if closestFood.X > myHead.X {
			preferredMoves["right"] = true
		}
		if closestFood.Y < myHead.Y {
			preferredMoves["down"] = true
		}
		if closestFood.Y > myHead.Y {
			preferredMoves["up"] = true
		}
		//fmt.Println("Preferred Moves:", preferredMoves)
	}

	return preferredMoves
}

func FindHazardDangerMoves(state GameState, dangerMoves map[string]bool) map[string]bool {
	myHead := state.You.Body[0]
	hazards := state.Board.Hazards

	if len(hazards) > 0 {
		for i := 0; i < len(hazards); i++ {
			if (myHead.X-1 == hazards[i].X) && (myHead.Y == hazards[i].Y) {
				dangerMoves["left"] = true
			}
			if (myHead.X+1 == hazards[i].X) && (myHead.Y == hazards[i].Y) {
				dangerMoves["right"] = true
			}
			if (myHead.X == hazards[i].X) && (myHead.Y-1 == hazards[i].Y) {
				dangerMoves["down"] = true
			}
			if (myHead.X == hazards[i].X) && (myHead.Y+1 == hazards[i].Y) {
				dangerMoves["up"] = true
			}
		}
	}

	return dangerMoves
}
