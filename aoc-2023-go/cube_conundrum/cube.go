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

type allCubes struct {
	cubes map[CubeType]string
}

var ALL_CUBES_VALUES = allCubes{
	cubes: map[CubeType]string{
		Red:   "red",
		Green: "green",
		Blue:  "blue",
	},
}

func newCube(cubeTypeData string) cube {
	return cube{
		cType: CubeType(cubeTypeData),
	}
}

func newCubeTypeStringToIntMap() map[CubeType]int {
	return map[CubeType]int{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
}
