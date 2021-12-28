package main

import (
	"log"
	"math/rand"
)

func ReturnChosenMove(state GameState, safeMoves []string, safeMovesTowardFood []string, largestLookDirection string) string {
	var nextMove string

	// if there are no safeMoves whatsoever, move down (ABORT MISSION)
	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		// else, some safe moves exist
		// if there are safe moves toward food, move that way
		if len(safeMovesTowardFood) > 0 {
			//fmt.Println("Choosing safeMoveTowardFood!")
			nextMove = safeMovesTowardFood[rand.Intn(len(safeMovesTowardFood))]
		} else {
			// if there are no safe moves toward food, just choose a random move from safe moves
			//fmt.Println("Choosing any safe move!")
			if contains(safeMoves, largestLookDirection) {
				nextMove = largestLookDirection
			} else {
				nextMove = safeMoves[rand.Intn(len(safeMoves))]
			}

		}
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return nextMove
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
