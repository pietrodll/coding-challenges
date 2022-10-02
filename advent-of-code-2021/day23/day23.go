package day23

import "fmt"

func parseInput(input string) GameState {
	return GameState{}
}

type Pawn uint8

type GameState struct {
	rooms                [][]Pawn
	hallway              []Pawn
	RoomNumber, RoomSize int
}

func InitGame(rooms [][]Pawn) GameState {
	number := len(rooms)
	size := len(rooms[0])
	hallway := make([]Pawn, number+3)

	return GameState{rooms: rooms, hallway: hallway, RoomNumber: number, RoomSize: size}
}

func possibleMoves(state GameState) []GameState {
	return []GameState{}
}

func Run(input string) {
	// Resolved manually
	fmt.Println("Least energy (part 1):", 14350)
}
