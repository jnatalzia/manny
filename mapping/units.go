package mapping

import (
	"../types"
	"../utils"
)

var currentUnits []*types.Civilian
var logger = utils.GetLogger()

func BuildCivs() {
	locOptions := []types.UnitLocationID{
		"1",
		"2",
		"3",
		"4",
		"5",
	}
	for i := 0; i < 10; i++ {
		currentUnits = append(currentUnits, &types.Civilian{
			IsAI:     true,
			Location: utils.GetRandomLocID(locOptions),
		})
	}

	logger.Print(currentUnits)
}

func GetCivs() []*types.Civilian {
	return currentUnits
}
