package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMyCubeConundrumPosibleGames(t *testing.T) {

	fmt.Println("Testing cube conundrum for possible games!")

	filePaths := []string{"cube_conundrum_one.txt", "cube_conundrum_data.txt"}
	maps := []bag{
		{
			bagContents: map[CubeType]int{
				Red:   12,
				Green: 13,
				Blue:  14,
			},
		},
		{
			bagContents: map[CubeType]int{
				Red:   12,
				Green: 13,
				Blue:  14,
			},
		},
	}
	results := []int{8, 2076}

	for i, path := range filePaths {
		startTime := time.Now()

		result := cubeConundrumPosibleGames(path, maps[i])

		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)

		if result != results[i] {
			t.Errorf("expected: %v, but got %v, for file: %s", results[i], result, path)
		} else {
			fmt.Printf("Test case passed for file: %s, time: %vms \n", path, elapsedTime.Milliseconds())
		}
	}

	fmt.Println()
}

func TestMyCubeConundrumMinimumSetOfGames(t *testing.T) {

	fmt.Println("Testing cube conundrum for power of minimum set games!")
	
	filePaths := []string{"cube_conundrum_one.txt", "cube_conundrum_data.txt"}
	results := []int{2286, 70950}

	for i, path := range filePaths {
		startTime := time.Now()

		result := cubeConundrumMinimumSetGames(path)

		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)

		if result != results[i] {
			t.Errorf("expected: %v, but got %v, for file: %s", results[i], result, path)
		} else {
			fmt.Printf("Test case passed for file: %s, time: %vms \n", path, elapsedTime.Milliseconds())
		}
	}
	fmt.Println()
}
