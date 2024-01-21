package main

import (
	"strconv"
	"strings"
)

type game struct {
	id     int
	rounds []round
	power  int
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

	myGame := game{
		id:     gameId,
		rounds: roundsPlayed,
	}
	return myGame
}

func (g *game) isPosible(b bag) bool {

	isPosible := true
	for _, round := range g.rounds {

		for _, c := range ALL_CUBES_VALUES.cubes {
			if !round.isValid(CubeType(c), b.bagContents[CubeType(c)]) {
				isPosible = false
				break
			}
		}

	}
	return isPosible
}

func (g *game) calcPower() {

	power := 1
	highestPlayed := newCubeTypeStringToIntMap()
	for _, r := range g.rounds {

		r.calcHighestNumberOfCubePlayed()
		for _, c := range ALL_CUBES_VALUES.cubes {
			if r.highestCubePlayed[CubeType(c)] > highestPlayed[CubeType(c)] {
				highestPlayed[CubeType(c)] = r.highestCubePlayed[CubeType(c)]
			}
		}

	}

	for _, countCubeType := range highestPlayed {
		power *= countCubeType
	}

	g.power = power
}
