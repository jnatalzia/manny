package gameplay

import (
	"../mapping"
	"../types"
	"../utils"
)

var logger = utils.GetLogger()

func GenerateIntents(civs []*types.Civilian) {
	for _, civ := range civs {
		civ.MoveIntent = utils.GetRandomLocID(mapping.LocIDs)
	}
	logger.Print("Intents Generated")
}

func MoveUnits(civs []*types.Civilian) {
	for i, civ := range civs {
		civ.Location = civ.MoveIntent
		logger.Printf("Civilian %d moved. %s", i, civ.Location)
	}
}
