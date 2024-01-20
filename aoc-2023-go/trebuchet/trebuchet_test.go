package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMyTrebuchet(t *testing.T) {

	filePaths := []string{"trebuchet_data_one.txt", "trebuchet_data_two.txt", "./trebuchet_data.txt"}
	results := []int{142, 281, 54249}

	for i, path := range filePaths {
		startTime := time.Now()

        result := trebuchet(path)

        endTime := time.Now()
        elapsedTime := endTime.Sub(startTime)

		if result != results[i] {
			t.Errorf("expected: %v, but got %v, for file: %s", results[i], result, path)
		} else {
			fmt.Printf("Test case passed for file: %s, time: %vms \n", path, elapsedTime.Milliseconds())
		}
	}
}
