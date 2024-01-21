package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type bag struct {
	bagContents map[CubeType]int
}

/*
*
Given the contents of the bag that the elf is carrying, find out for each game played
if it would have been possible!
*
*/
func cubeConundrumPosibleGames(filePath string, contentsOfBag bag) int {
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

/*
*
Find out for each game played, the minimum cubes that would be required.
For each game find out its "power" which is the minimum cubes that would be required multiplied with each other
Find the sum of the powers for all the games!
*
*/
func cubeConundrumMinimumSetGames(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	myTournament := newTournamentMin()
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		g := newGame(text)
		myTournament.addGame(g)
	}

	return myTournament.totalSumOfMinimumSetPower()
}

func main() {
	bagContents := bag{
		bagContents: map[CubeType]int{
			Red:   12,
			Green: 13,
			Blue:  14,
		},
	}

	fmt.Println("------------------------------------")
	fmt.Println("Total sum of possible Game ids:", cubeConundrumPosibleGames("./cube_conundrum_one.txt", bagContents))
	fmt.Println("------------------------------------")

	fmt.Println("------------------------------------")
	fmt.Println("Total sum of power of games", cubeConundrumMinimumSetGames("./cube_conundrum_one.txt"))
	fmt.Println("------------------------------------")

}
