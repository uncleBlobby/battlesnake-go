package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"fmt"
	"log"
	"math/rand"
)

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
	possibleMoves = AvoidMyNeck(state, possibleMoves)

	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	possibleMoves = AvoidWalls(state, possibleMoves)

	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
	possibleMoves = AvoidSelf(state, possibleMoves)

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.
	possibleMoves = AvoidOtherSnakes(state, possibleMoves)
	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.
	preferredMoves = PreferFoodMoves(state, preferredMoves)

	// TODO: avoid hazard sauce when possible!
	dangerMoves = FindHazardDangerMoves(state, dangerMoves)
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
