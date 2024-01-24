package main

import "unicode"

type PartNumber struct {
	number int
	start  int
	end    int
}

func newPartNumber(number, start, end int) *PartNumber {
	return &PartNumber{
		number: number,
		start:  start,
		end:    end,
	}
}

func (pn *PartNumber) isSymbolRightOfNumber(line string) bool {
	for i, char := range line {
		if unicode.IsDigit(char) || char == '.' {
			continue
		}

		if pn.end == i {
			return true
		}
	}
	return false
}

func (pn *PartNumber) isSymbolLeftOfNumber(line string) bool {
	for i, char := range line {
		if unicode.IsDigit(char) || char == '.' {
			continue
		}

		if pn.start-1 == i {
			return true
		}
	}
	return false
}

func (pn *PartNumber) isSymbolTopOfNumber(line string) bool {
	for i, char := range line {
		if unicode.IsDigit(char) || char == '.' {
			continue
		}

		if pn.start <= i && pn.end-1 >= i {
			return true
		}
	}
	return false
}

//its the same but the line changes, keeping it for readability
func (pn *PartNumber) isSymbolBottomOfNumber(line string) bool {
	return pn.isSymbolTopOfNumber(line)
}
