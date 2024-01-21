package main

type tournament struct {
	games         []game
	tBag bag
}

func newTournament(tBag bag) tournament {
	return tournament{
		tBag: tBag,
	}
}

func (t *tournament) addGame(myGame game) {
	t.games = append(t.games, myGame)
}

func (t *tournament) totalSumOfPossibleGames() int {

	sumOfGameIds := 0
	for _, g := range t.games {
		if g.isPosible(t.tBag) {
			sumOfGameIds = sumOfGameIds + g.id
		}
	}

	return sumOfGameIds
}
