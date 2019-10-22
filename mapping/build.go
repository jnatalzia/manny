package mapping

import "../types"

var GameMap types.Map = map[types.UnitLocationID]types.Location{
	"1": types.Location{
		ID:        "1",
		Neighbors: []types.UnitLocationID{"2", "4", "5"},
	},
	"2": types.Location{
		ID:        "2",
		Neighbors: []types.UnitLocationID{"1"},
	},
	"3": types.Location{
		ID:        "3",
		Neighbors: []types.UnitLocationID{"5", "6"},
	},
	"4": types.Location{
		ID:        "4",
		Neighbors: []types.UnitLocationID{"1"},
	},
	"5": types.Location{
		ID:        "5",
		Neighbors: []types.UnitLocationID{"3", "1"},
	},
	"6": types.Location{
		ID:        "6",
		Neighbors: []types.UnitLocationID{"3"},
	},
}

var LocIDs []types.UnitLocationID

func InitHelperData() {
	for k := range GameMap {
		LocIDs = append(LocIDs, k)
	}
}
