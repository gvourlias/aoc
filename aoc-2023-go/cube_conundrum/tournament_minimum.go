package main

type tournamentMinimum struct {
	games         []game
}

func newTournamentMin() tournamentMinimum {
	return tournamentMinimum{}
}

func (t *tournamentMinimum) addGame(myGame game) {
	t.games = append(t.games, myGame)
}

func (t *tournamentMinimum) totalSumOfMinimumSetPower() int {
	totalPower := 0
	for _, g := range t.games {
		g.calcPower()
		totalPower = totalPower + g.power
	}

	return totalPower
}
