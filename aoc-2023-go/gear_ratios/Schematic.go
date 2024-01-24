package main

type Schematic struct {
	partLines       []string
	schematicChunks []SchematicChunk
}

func newSchematic() *Schematic {
	return &Schematic{}
}

func (s *Schematic) addLine(partLineData string) {
	s.partLines = append(s.partLines, partLineData)
}

// all chunks are 3 lines except for the first and last which are two
func (s *Schematic) createSchematicChunks() {

	lengthOfPartLines := len(s.partLines)

	//for the first lines take the first two and create a chunk
	s.schematicChunks = append(s.schematicChunks, *newSchematicChunk().theFirstChunk(true).createPartLines([]string{s.partLines[0], s.partLines[1]}))

	//for the last, take the last two lines and create a chunk
	s.schematicChunks = append(s.schematicChunks, *newSchematicChunk().theLastChunk(true).createPartLines([]string{s.partLines[lengthOfPartLines-2], s.partLines[lengthOfPartLines-1]}))

	for i := 1; i < lengthOfPartLines - 1; i++ {
		s.schematicChunks = append(s.schematicChunks, *newSchematicChunk().createPartLines([]string{s.partLines[i-1], s.partLines[i], s.partLines[i+1]}))
	}
}

func (s *Schematic) sumOfValidPartNumbers() int {

	sum := 0

	for _, sc := range s.schematicChunks {
		sc.findValidPartNumbers()
		for _, vpn := range sc.validNumbers {
			sum = sum + vpn.number
		}
	}

	return sum
}
