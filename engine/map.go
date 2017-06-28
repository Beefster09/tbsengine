package engine

// GameMap represents a map in game
type GameMap struct {
	name         string
	terrain      [][]int8
	terrainIndex []Terrain
}

func makeMap(name string, width int, height int) GameMap {
	buffer := make([]int8, width*height)
	slice := make([][]int8, height)
	for i := 0; i < height; i++ {
		slice[i] = buffer[width*i : width*(i+1)]
	}
	return GameMap{name, slice, nil}
}
