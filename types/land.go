package types

type UnitLocationID string

type Location struct {
	ID        UnitLocationID
	Neighbors []UnitLocationID
}

type Map map[UnitLocationID]Location
