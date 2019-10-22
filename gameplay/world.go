package gameplay

import (
	"../utils"
)

var currentUnits []*Civilian
var logger = utils.GetLogger()

func BuildCivs() {
	// locOptions := []string{
	// 	"1",
	// 	"2",
	// 	"3",
	// 	"4",
	// 	"5",
	// }
	for i := 0; i < 1; i++ { //10; i++ {
		currentUnits = append(currentUnits, &Civilian{
			IsAI:       true,
			Location:   "1",
			MoveIntent: "6",
		})
	}

	logger.Print(currentUnits)
}

func GetCivs() []*Civilian {
	return currentUnits
}
