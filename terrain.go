package tbsengine

// Terrain represents a type of square on a map
type Terrain struct {
	name       string
	attributes map[string]any
}

// MovementClass represents how well a unit can move through a terrain type
// If a terrain type is missing from the costs map, it means that terrain is not crossable by that movement class
type MovementClass struct {
	name  string
	costs map[string]int
}
