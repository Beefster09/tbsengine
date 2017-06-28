package engine

type Unit struct {
	name      string
	hp        int
	maxHP     int
	moveClass MovementClass
	moveSpeed int
	abilities []AbilityRule
}
