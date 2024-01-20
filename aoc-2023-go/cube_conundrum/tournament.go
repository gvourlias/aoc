package main

type tournament struct {
	games         []game
	contentsOfBag map[CubeType]int
}

func newTournament(contentsOfBag map[CubeType]int) tournament {
	return tournament{
		contentsOfBag: contentsOfBag,
	}
}

func (t *tournament) addGame(myGame game) {
	t.games = append(t.games, myGame)
}

func (t *tournament) totalSumOfPossibleGames() int {

	sumOfGameIds := 0
	for _, g := range t.games {
		if g.isPosible(t.contentsOfBag) {
			sumOfGameIds = sumOfGameIds + g.id
		}
	}

	return sumOfGameIds
}
