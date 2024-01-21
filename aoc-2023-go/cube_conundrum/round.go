package main

import (
	"strings"
)

type round struct {
	handsPlayed       []hand
	highestCubePlayed map[CubeType]int
}

func newRound(roundData string) round {
	hands := strings.Split(roundData, ",")

	handsPlayed := []hand{}
	for _, handData := range hands {
		hand := newHand(handData)
		handsPlayed = append(handsPlayed, hand)
	}

	myRound := round{
		handsPlayed: handsPlayed,
	}
	return myRound
}

func (r *round) isValid(cubeType CubeType, upperLimitOfCubeTypePlayed int) bool {
	isValid := true

	for _, handPlayed := range r.handsPlayed {

		if handPlayed.cubePlayed.cType != cubeType {
			continue
		}

		if handPlayed.numberPlayed > upperLimitOfCubeTypePlayed {
			isValid = false
			break
		}

	}
	return isValid
}

func (r *round) calcHighestNumberOfCubePlayed() {

	highestCubePlayed := newCubeTypeStringToIntMap()

	for _, hand := range r.handsPlayed {

		for _, c := range ALL_CUBES_VALUES.cubes {
			if hand.cubePlayed.cType == CubeType(c) && hand.numberPlayed > highestCubePlayed[CubeType(c)] {
				highestCubePlayed[CubeType(c)] = hand.numberPlayed
			}
		}
		
	}

	r.highestCubePlayed = highestCubePlayed
}
