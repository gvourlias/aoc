package main

import (
	"fmt"
)

type SchematicChunk struct {
	isFirstChunk bool
	isLastChunk  bool
	partLines    []*PartLine
	validNumbers []*PartNumber
}

func newSchematicChunk() *SchematicChunk {
	return &SchematicChunk{}
}

func (sc *SchematicChunk) theFirstChunk(isFirstChunk bool) *SchematicChunk {
	sc.isFirstChunk = isFirstChunk
	return sc
}

func (sc *SchematicChunk) theLastChunk(isLastChunk bool) *SchematicChunk {
	sc.isLastChunk = isLastChunk
	return sc
}

func (sc *SchematicChunk) createPartLines(partLineData []string) *SchematicChunk {
	var scPartLines []*PartLine
	for _, singlePartLineData := range partLineData {
		partLine := newPartLine(singlePartLineData)
		partLine.calculatePartNumbers()
		scPartLines = append(scPartLines, partLine)
	}
	sc.partLines = scPartLines

	return sc
}

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..

func (sc *SchematicChunk) findValidPartNumbers() {

	if sc.isFirstChunk {
		firstLine := *sc.partLines[0]
		secondLine := *sc.partLines[1]

		if len(firstLine.partNumbers) == 0 {
			return
		}

		for _, partNumber := range firstLine.partNumbers {
			isRight := partNumber.isSymbolRightOfNumber(firstLine.line)
			isLeft := partNumber.isSymbolLeftOfNumber(firstLine.line)
			isRightDiagonal := partNumber.isSymbolRightOfNumber(secondLine.line)
			isLeftDiagonal := partNumber.isSymbolLeftOfNumber(secondLine.line)
			isBottom := partNumber.isSymbolBottomOfNumber(secondLine.line)
			if isRight ||
				isLeft ||
				isRightDiagonal ||
				isLeftDiagonal ||
				isBottom {
				sc.validNumbers = append(sc.validNumbers, partNumber)
			}
		}

	} else if sc.isLastChunk {
		secondToLast := *sc.partLines[0]
		lastLine := *sc.partLines[1]

		if len(lastLine.partNumbers) == 0 {
			return
		}

		for _, partNumber := range lastLine.partNumbers {
			isRight := partNumber.isSymbolRightOfNumber(lastLine.line)
			isLeft := partNumber.isSymbolLeftOfNumber(lastLine.line)
			isRightDiagonal := partNumber.isSymbolRightOfNumber(secondToLast.line)
			isLeftDiagonal := partNumber.isSymbolLeftOfNumber(secondToLast.line)
			isTop := partNumber.isSymbolTopOfNumber(secondToLast.line)
			if isRight ||
				isLeft ||
				isRightDiagonal ||
				isLeftDiagonal ||
				isTop {
				sc.validNumbers = append(sc.validNumbers, partNumber)
			}
		}

	} else {
// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
		firstLine := *sc.partLines[0]
		secondLine := *sc.partLines[1]
		thirdLine := *sc.partLines[2]

		if len(secondLine.partNumbers) == 0 {
			return
		}

		for _, partNumber := range secondLine.partNumbers {
			isRight := partNumber.isSymbolRightOfNumber(secondLine.line)
			isLeft := partNumber.isSymbolLeftOfNumber(secondLine.line)
			isRightDiagonal_TopLine := partNumber.isSymbolRightOfNumber(firstLine.line)
			isLeftDiagonal_TopLine := partNumber.isSymbolLeftOfNumber(firstLine.line)
			isTop := partNumber.isSymbolTopOfNumber(firstLine.line)
			isRightDiagonal_BottomLine := partNumber.isSymbolRightOfNumber(thirdLine.line)
			isLeftDiagonal_BottomLine := partNumber.isSymbolLeftOfNumber(thirdLine.line)
			isBottom := partNumber.isSymbolBottomOfNumber(thirdLine.line)

			if  isRight ||
			isLeft ||
			isRightDiagonal_TopLine ||
			isLeftDiagonal_TopLine ||
			isTop ||
			isRightDiagonal_BottomLine ||
			isLeftDiagonal_BottomLine ||
			isBottom {
				sc.validNumbers = append(sc.validNumbers, partNumber)
			}
		}
	}

	for _, validNum := range sc.validNumbers {
		fmt.Println(*validNum)
	}
}
