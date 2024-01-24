package main

import (
	"bufio"
	"log"
	"os"
)

func calcAllPartsOfEngineSchematics(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("failed to close file")
		}
	}(file)

	scanner := bufio.NewScanner(file)

	mySchematic := newSchematic()
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		mySchematic.addLine(text)
	}

	mySchematic.createSchematicChunks()
	return mySchematic.sumOfValidPartNumbers()
}

func main() {
	calcAllPartsOfEngineSchematics("./data/gear_ratios_one.txt")
}
