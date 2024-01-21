package main

import (
	"strconv"
	"strings"
)

type game struct {
	id     int
	rounds []round
}

func newGame(gameData string) game {
	gamePortionParts := strings.Split(gameData, ":")
	gamePortionPart := gamePortionParts[0]
	gameIdParts := strings.Split(gamePortionPart, " ")
	gameIdString := gameIdParts[1]

	gameId, err := strconv.Atoi(gameIdString)
	if err != nil {
		panic(err)
	}

	roundsData := strings.Split(gamePortionParts[1], ";")

	roundsPlayed := []round{}
	for _, roundData := range roundsData {
		myRound := newRound(roundData)
		roundsPlayed = append(roundsPlayed, myRound)
	}

	return game{
		id:     gameId,
		rounds: roundsPlayed,
	}
}

func (g *game) isPosible(contentsOfBag map[CubeType]int) bool {

	isPosible := true
	for _, round := range g.rounds {
		if !round.isValid(Red, contentsOfBag[Red]) ||
			!round.isValid(Blue, contentsOfBag[Blue]) ||
			!round.isValid(Green, contentsOfBag[Green]) {
			isPosible = false
			break
		}
	}
	return isPosible
}
