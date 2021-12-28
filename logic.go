package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

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

func avoidMyNeck(state GameState, possibleMoves map[string]bool) map[string]bool {
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

func avoidWalls(state GameState, possibleMoves map[string]bool) map[string]bool {
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

func avoidSelf(state GameState, possibleMoves map[string]bool) map[string]bool {
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

func avoidOtherSnakes(state GameState, possibleMoves map[string]bool) map[string]bool {
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

func preferFoodMoves(state GameState, preferredMoves map[string]bool) map[string]bool {
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

func findHazardDangerMoves(state GameState, dangerMoves map[string]bool) map[string]bool {
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

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "uncleBlobby", // TODO: Your Battlesnake username
		Color:      "#888888",     // TODO: Personalize
		Head:       "default",     // TODO: Personalize
		Tail:       "default",     // TODO: Personalize
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	preferredMoves := map[string]bool{
		"up":    false,
		"down":  false,
		"left":  false,
		"right": false,
	}

	dangerMoves := map[string]bool{
		"up":    false,
		"down":  false,
		"left":  false,
		"right": false,
	}

	// Step 0: Don't let your Battlesnake move back in on it's own neck
	possibleMoves = avoidMyNeck(state, possibleMoves)

	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	possibleMoves = avoidWalls(state, possibleMoves)

	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
	possibleMoves = avoidSelf(state, possibleMoves)

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.
	possibleMoves = avoidOtherSnakes(state, possibleMoves)
	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.
	preferredMoves = preferFoodMoves(state, preferredMoves)

	// TODO: avoid hazard sauce when possible!
	dangerMoves = findHazardDangerMoves(state, dangerMoves)
	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	hazardMoves := []string{}
	for move, isDangerous := range dangerMoves {
		if isDangerous {
			hazardMoves = append(hazardMoves, move)
		}
	}

	foodMoves := []string{}
	for move, isPreferred := range preferredMoves {
		if isPreferred {
			foodMoves = append(foodMoves, move)
		}
	}

	//fmt.Println("foodMoves:", foodMoves)

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	// if safeMoves is greater in length than hazardMoves, remove hazardMoves from safeMoves
	fmt.Println("safeMoves before haz reduction:", safeMoves)
	if len(safeMoves) > len(hazardMoves) {
		for i := 0; i < len(hazardMoves); i++ {
			for j := 0; j < len(safeMoves); j++ {
				if safeMoves[j] == hazardMoves[i] {
					safeMoves = append(safeMoves[:j], safeMoves[j+1:]...)

				}
			}
		}
	}
	// safeMoves should now have hazardMoves removed!
	fmt.Println("safeMoves after haz reduction:", safeMoves)
	fmt.Println("hazardMoves: ", hazardMoves)

	//fmt.Println("safeMoves:", safeMoves)

	safeMovesTowardFood := []string{}
	for i := 0; i < len(safeMoves); i++ {
		for j := 0; j < len(foodMoves); j++ {
			if safeMoves[i] == foodMoves[j] {
				safeMovesTowardFood = append(safeMovesTowardFood, foodMoves[j])
			}
		}
	}

	//fmt.Println("safeMovesTowardFood:", safeMovesTowardFood)
	//fmt.Println("len(safeMovesTowardFood):", len(safeMovesTowardFood))
	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		if len(safeMovesTowardFood) > 0 {
			//fmt.Println("Choosing safeMoveTowardFood!")
			nextMove = safeMovesTowardFood[rand.Intn(len(safeMovesTowardFood))]
		} else {
			//fmt.Println("Choosing any safe move!")
			nextMove = safeMoves[rand.Intn(len(safeMoves))]
		}
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
