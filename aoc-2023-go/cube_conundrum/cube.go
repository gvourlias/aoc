package main

type cube struct {
	cType CubeType
}

type CubeType string

const (
	Red   CubeType = "red"
	Green CubeType = "green"
	Blue  CubeType = "blue"
)

func newCube(cubeTypeData string) cube {
	return cube{
		cType: CubeType(cubeTypeData),
	}
}