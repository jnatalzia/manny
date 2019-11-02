package mapping

import "../types"

var GameMap types.Map = map[string]types.Location{
	"0": types.Location{
		ID:        "0",
		Neighbors: []string{"2"},
	},
	"1": types.Location{
		ID:        "1",
		Neighbors: []string{"2", "5", "6"},
	},
	"2": types.Location{
		ID:        "2",
		Neighbors: []string{"0", "1", "3", "6"},
	},
	"3": types.Location{
		ID:        "3",
		Neighbors: []string{"2", "7"},
	},
	"4": types.Location{
		ID:        "4",
		Neighbors: []string{"5"},
	},
	"5": types.Location{
		ID:        "5",
		Neighbors: []string{"1", "4", "6", "9"},
	},
	"6": types.Location{
		ID:        "6",
		Neighbors: []string{"2", "5", "7", "10"},
	},
	"7": types.Location{
		ID:        "7",
		Neighbors: []string{"3", "6", "8", "11"},
	},
	"8": types.Location{
		ID:        "8",
		Neighbors: []string{"7"},
	},
	"9": types.Location{
		ID:        "9",
		Neighbors: []string{"5", "10"},
	},
	"10": types.Location{
		ID:        "10",
		Neighbors: []string{"6", "9", "11", "12"},
	},
	"11": types.Location{
		ID:        "11",
		Neighbors: []string{"10", "7"},
	},
	"12": types.Location{
		ID:        "12",
		Neighbors: []string{"10"},
	},
}
