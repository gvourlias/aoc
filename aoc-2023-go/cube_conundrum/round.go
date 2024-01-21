package main

import (
	"strings"
)

type round struct {
	handsPlayed []hand
}

func newRound(roundData string) round {
	hands := strings.Split(roundData, ",")

	handsPlayed := []hand{}
	for _, handData := range hands {
		hand := newHand(handData)
		handsPlayed = append(handsPlayed, hand)
	}

	return round{
		handsPlayed: handsPlayed,
	}
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
