package main

import (
	"unicode"
)

type PartLine struct {
	line        string
	partNumbers []*PartNumber
}

func newPartLine(lineData string) *PartLine {
	return &PartLine{line: lineData}
}

func (pl *PartLine) calculatePartNumbers() {
	var partNumbersFound []*PartNumber
	digitsFound := []int{}
	foundDigit := false
	startPositionOfPartNumber := 0

	for i, char := range pl.line {
		if unicode.IsDigit(char) {
			if !foundDigit {
				startPositionOfPartNumber = i
			}
			foundDigit = true

			digitsFound = append(digitsFound, int(char-'0'))
		} else {
			if foundDigit {
				partNumbersFound = append(partNumbersFound, newPartNumber(calculateResultOfDigits(digitsFound), startPositionOfPartNumber, i))
				digitsFound = []int{}
			}
			foundDigit = false
		}
	}

	//if number was found right at the edge of the line
	if len(digitsFound) > 0 {
		partNumbersFound = append(partNumbersFound, newPartNumber(calculateResultOfDigits(digitsFound), startPositionOfPartNumber, len(pl.line)))
	}

	pl.partNumbers = partNumbersFound
}

func calculateResultOfDigits (digits []int) int{
	result := 0
	for _, j := range digits {
		result = result*10 + j
	}

	return result;
}
