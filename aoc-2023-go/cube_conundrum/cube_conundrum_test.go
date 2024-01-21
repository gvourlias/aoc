package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMyCubeConundrum(t *testing.T) {

	filePaths := []string{"cube_conundrum_one.txt", "cube_conundrum_data.txt"}
	maps := []map[CubeType]int{
		{
			Red:   12,
			Green: 13,
			Blue:  14,
		},
		{
			Red:   12,
			Green: 13,
			Blue:  14,
		},
		{
			Red:   10,
			Green: 10,
			Blue:  10,
		},
	}
	results := []int{8, 2076}

	for i, path := range filePaths {
		startTime := time.Now()

		result := cubeConundrum(path, maps[i])

		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)

		if result != results[i] {
			t.Errorf("expected: %v, but got %v, for file: %s", results[i], result, path)
		} else {
			fmt.Printf("Test case passed for file: %s, time: %vms \n", path, elapsedTime.Milliseconds())
		}
	}
}
