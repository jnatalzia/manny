package gameplay

import (
	"../utils"
)

var currentUnits []*Civilian
var logger = utils.Logger
var locOptions = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"11",
	"12",
}

func BuildCivs() {
	for i := 0; i < 10; i++ {
		currentUnits = append(currentUnits, &Civilian{
			ID:           utils.GetGUID(),
			IsAI:         true,
			Location:     utils.GetRandomLocID(locOptions),
			MoveIntent:   utils.GetRandomLocID(locOptions),
			PathPosition: 0,
		})
	}

	logger.Print(currentUnits)
}

func GetCivs() []*Civilian {
	return currentUnits
}

func TakeTurn() {
	for _, c := range currentUnits {
		c.Location = c.Path[c.PathPosition]
		c.PathPosition++

		if c.MoveIntent == c.Location {
			logger.Printf("Civilian arrived at their destination %s\n", c.MoveIntent)
			logger.Println("Making new intent for civ")
			c.PathPosition = 0
			c.MoveIntent = utils.GetRandomLocID(locOptions)
			GeneratePathing([]*Civilian{c})
		} else {
			logger.Printf("Civilian moved to %s\n", c.Path[c.PathPosition-1])
		}
	}
}
