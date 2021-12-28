package main

import (
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

// TODO: More GameState test cases!
