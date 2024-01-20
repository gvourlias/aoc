package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func cubeConundrum(filePath string, contentsOfBag map[CubeType]int) int {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	myTournament := newTournament(contentsOfBag)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		g := newGame(text)
		myTournament.addGame(g)
	}

	return myTournament.totalSumOfPossibleGames()
}

func main() {
	fmt.Println("------------------------------------")
	fmt.Println("Total sum of possible Game ids:", cubeConundrum("./cube_conundrum_one.txt", map[CubeType]int{
		Red:   12,
		Green: 13,
		Blue:  14,
	}))
	fmt.Println("------------------------------------")
}
