package main

import (
	"fmt"
	"testing"
)

func TestNeckAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 2, Y: 0},
		Body: []Coord{{X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move == "left" {
			t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
		}
	}
}

func TestRightWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right wall
		Head: Coord{X: 10, Y: 5},
		Body: []Coord{{X: 10, Y: 5}, {X: 10, Y: 4}, {X: 10, Y: 3}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move right
		if nextMove.Move == "right" {
			t.Errorf("snake moved into the right wall, %s", nextMove.Move)
		}
	}
}

func TestLeftWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing left wall
		Head: Coord{X: 0, Y: 5},
		Body: []Coord{{X: 0, Y: 5}, {X: 0, Y: 4}, {X: 0, Y: 3}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move == "left" {
			t.Errorf("snake moved into the left wall, %s", nextMove.Move)
		}
	}
}

func TestBottomWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing bottom wall
		Head: Coord{X: 5, Y: 0},
		Body: []Coord{{X: 5, Y: 0}, {X: 5, Y: 1}, {X: 5, Y: 2}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move down
		if nextMove.Move == "down" {
			t.Errorf("snake moved into the bottom wall, %s", nextMove.Move)
		}
	}
}

func TestTopWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing top wall
		Head: Coord{X: 5, Y: 10},
		Body: []Coord{{X: 5, Y: 10}, {X: 5, Y: 9}, {X: 5, Y: 8}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move up
		if nextMove.Move == "up" {
			t.Errorf("snake moved into the top wall, %s", nextMove.Move)
		}
	}
}

func TestSelfAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		//Length 6, coiled up in a knot.
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 4, Y: 5}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 6, Y: 5}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Never Move anywhere but up
		if nextMove.Move != "up" {
			t.Errorf("snake moved into itself, %s", nextMove.Move)
		}
	}
}

func TestOtherSnakeAvoidance1(t *testing.T) {
	// Two Snakes beside, neck to the south
	me := Battlesnake{
		//Length 3, middle of board next to other snakes
		Head:   Coord{X: 5, Y: 5},
		Body:   []Coord{{X: 5, Y: 5}, {X: 5, Y: 4}, {X: 5, Y: 3}},
		Length: 3,
	}
	enemy1 := Battlesnake{
		Head:   Coord{X: 4, Y: 5},
		Body:   []Coord{{X: 4, Y: 5}, {X: 4, Y: 4}, {X: 4, Y: 3}},
		Length: 3,
	}
	enemy2 := Battlesnake{
		Head:   Coord{X: 6, Y: 5},
		Body:   []Coord{{X: 6, Y: 5}, {X: 6, Y: 4}, {X: 6, Y: 3}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1, enemy2},
		},
		You: me,
	}

	// Act 1,000x
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Don't move Left or Right into other snakes!
		if nextMove.Move == "left" || nextMove.Move == "right" {
			t.Errorf("snake moved into enemy, %s", nextMove.Move)
		}
	}
}

func TestOtherSnakeAvoidance2(t *testing.T) {
	// Two Snakes beside, neck to the north
	me := Battlesnake{
		//Length 3, middle of board next to other snakes
		Head:   Coord{X: 5, Y: 5},
		Body:   []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
		Length: 3,
	}
	enemy1 := Battlesnake{
		Head:   Coord{X: 4, Y: 5},
		Body:   []Coord{{X: 4, Y: 5}, {X: 4, Y: 4}, {X: 4, Y: 3}},
		Length: 3,
	}
	enemy2 := Battlesnake{
		Head:   Coord{X: 6, Y: 5},
		Body:   []Coord{{X: 6, Y: 5}, {X: 6, Y: 4}, {X: 6, Y: 3}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1, enemy2},
		},
		You: me,
	}

	// Act 1,000x
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Don't move Left or Right into other snakes!
		if nextMove.Move == "left" || nextMove.Move == "right" {
			t.Errorf("snake moved into enemy, %s", nextMove.Move)
		}
	}
}

func TestMovesUpOrRightTowardFood(t *testing.T) {
	// Arrange
	me := Battlesnake{
		Head:   Coord{X: 5, Y: 5},
		Body:   []Coord{{X: 5, Y: 5}, {X: 5, Y: 4}, {X: 5, Y: 3}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Food:   []Coord{{X: 10, Y: 10}, {X: 9, Y: 9}},
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Should move Right or Up toward food!
		if (nextMove.Move != "right") && (nextMove.Move != "up") {
			t.Errorf("snake failed to move toward food, %s", nextMove.Move)
		}
	}

}

func TestMovesDownOrLeftTowardFood(t *testing.T) {
	// Arrange
	me := Battlesnake{
		Head:   Coord{X: 5, Y: 5},
		Body:   []Coord{{X: 5, Y: 5}, {X: 5, Y: 4}, {X: 5, Y: 3}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Food:   []Coord{{X: 0, Y: 0}, {X: 1, Y: 1}},
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Should move Right or Up toward food!
		if (nextMove.Move != "down") && (nextMove.Move != "left") {
			t.Errorf("snake failed to move toward food, %s", nextMove.Move)
		}
	}

}

func TestSimpleHazardAvoidanceOneDirection(t *testing.T) {
	// Arrange
	me := Battlesnake{
		Head:   Coord{X: 1, Y: 5},
		Body:   []Coord{{X: 1, Y: 5}, {X: 1, Y: 4}, {X: 1, Y: 3}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Food:   []Coord{{}},
			Snakes: []Battlesnake{me},
			Hazards: []Coord{
				{X: 0, Y: 10}, {X: 0, Y: 9}, {X: 0, Y: 8},
				{X: 0, Y: 7}, {X: 0, Y: 6}, {X: 0, Y: 5},
				{X: 0, Y: 4}, {X: 0, Y: 3}, {X: 0, Y: 2},
				{X: 0, Y: 1}},
		},
		You: me,
	}

	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Should NOT move left into hazard!
		if nextMove.Move == "left" {
			t.Errorf("snake moved %s into hazard", nextMove.Move)
		}
	}
}

func TestSimpleHazardAvoidanceTwoDirections(t *testing.T) {
	// Arrange
	me := Battlesnake{
		Head:   Coord{X: 1, Y: 9},
		Body:   []Coord{{X: 1, Y: 9}, {X: 1, Y: 8}, {X: 1, Y: 7}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Food:   []Coord{{}},
			Snakes: []Battlesnake{me},
			Hazards: []Coord{
				{X: 0, Y: 10}, {X: 0, Y: 9}, {X: 0, Y: 8},
				{X: 0, Y: 7}, {X: 0, Y: 6}, {X: 0, Y: 5},
				{X: 0, Y: 4}, {X: 0, Y: 3}, {X: 0, Y: 2},
				{X: 0, Y: 1},
				{X: 1, Y: 10}, {X: 2, Y: 10}, {X: 3, Y: 10},
				{X: 4, Y: 10}, {X: 5, Y: 10}, {X: 6, Y: 10},
				{X: 7, Y: 10}, {X: 8, Y: 10}, {X: 9, Y: 10},
				{X: 10, Y: 10}},
		},
		You: me,
	}

	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Should NOT move left or up into hazard!
		if (nextMove.Move == "left") || (nextMove.Move == "up") {
			t.Errorf("snake moved %s into hazard", nextMove.Move)
		}
	}
}

func TestLookDistance(t *testing.T) {
	fmt.Println("Look Distance Test")
	// Arrange
	me := Battlesnake{
		Head:   Coord{X: 5, Y: 5},
		Body:   []Coord{{X: 5, Y: 5}, {X: 4, Y: 5}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 6, Y: 5}},
		Length: 6,
	}
	enemy1 := Battlesnake{
		Head:   Coord{X: 5, Y: 8},
		Body:   []Coord{{X: 5, Y: 8}, {X: 4, Y: 8}, {X: 3, Y: 8}},
		Length: 3,
	}

	state := GameState{
		Board: Board{
			Width:   11,
			Height:  11,
			Food:    []Coord{{}},
			Snakes:  []Battlesnake{me, enemy1},
			Hazards: []Coord{{}},
		},
		You: me,
	}
	lookDistance := map[string]int{
		"up":    0,
		"down":  0,
		"left":  0,
		"right": 0,
	}
	for i := 0; i < 1000; i++ {
		lookDistance := ResetLookDistance(lookDistance)
		lookDistance = DetermineOpenSpaces(state, lookDistance)
		//fmt.Println("Look Distance: ", lookDistance)
		if lookDistance["up"] < lookDistance["down"] {
			t.Error("look distance not working properly ", lookDistance)
		}
	}
}

// TODO: More GameState test cases!
