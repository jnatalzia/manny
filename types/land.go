package types

type UnitLocationID string

type Location struct {
	ID        UnitLocationID
	Neighbors []string
}

type Map map[UnitLocationID]Location
