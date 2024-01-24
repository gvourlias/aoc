package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCalcAllPartsOfEngineSchematics(t *testing.T) {

	fmt.Println("Testing calc all part of engine schematics!")

	filePaths := []string{"./data/gear_ratios_one.txt", "./data/gear_ratios_data.txt"}
	results := []int{4361, 0 }

	for i, path := range filePaths {
		startTime := time.Now()

		result := calcAllPartsOfEngineSchematics(path)

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
