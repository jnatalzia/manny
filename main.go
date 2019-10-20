package main

import (
	"time"

	"./gameplay"
	"./mapping"
	"./utils"
)

func main() {
	utils.InitRand()
	mapping.InitHelperData()
	mapping.BuildCivs()

	// Game Loop
	for {
		gameplay.GenerateIntents(mapping.GetCivs())

		gameplay.MoveUnits(mapping.GetCivs())
		time.Sleep(time.Second * 10)
	}
}
