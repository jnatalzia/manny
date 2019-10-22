package mapping

import "../types"

var GameMap types.Map = map[string]types.Location{
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
		Neighbors: []string{"5", "6"},
	},
	"4": types.Location{
		ID:        "4",
		Neighbors: []string{"1", "6"},
	},
	"5": types.Location{
		ID:        "5",
		Neighbors: []string{"3", "1", "6"},
	},
	"6": types.Location{
		ID:        "6",
		Neighbors: []string{"3", "4", "5"},
	},
}

var LocIDs []string

func InitHelperData() {
	for k := range GameMap {
		LocIDs = append(LocIDs, k)
	}
}
