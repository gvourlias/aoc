package main

import (
	"strconv"
	"strings"
)

type hand struct {
	numberPlayed int
	cubePlayed   cube
}

func newHand(handData string) hand {

	handData = strings.TrimSpace(handData)
	handDataSplit := strings.Split(handData, " ")

	numberPlayed, err := strconv.Atoi(handDataSplit[0])
	if err != nil {
		panic(err)
	}

	cube := newCube(handDataSplit[1])

	return hand{
		numberPlayed: numberPlayed,
		cubePlayed:   cube,
	}
}

func (h *hand) getCubePlayed() cube {
	return h.cubePlayed
}
