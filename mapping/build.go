package mapping

import "../types"

var GameMap types.Map = map[types.UnitLocationID]types.Location{
	"1": types.Location{
		ID:        "1",
		Neighbors: []string{"2", "4", "5"},
	},
	"2": types.Location{
		ID:        "2",
		Neighbors: []string{"1"},
	},
	"3": types.Location{
		ID:        "3",
		Neighbors: []string{"5"},
	},
	"4": types.Location{
		ID:        "4",
		Neighbors: []string{"1"},
	},
	"5": types.Location{
		ID:        "5",
		Neighbors: []string{"3", "1"},
	},
}

var LocIDs []types.UnitLocationID

func InitHelperData() {
	for k := range GameMap {
		LocIDs = append(LocIDs, k)
	}
}
