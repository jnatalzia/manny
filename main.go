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
	gameplay.BuildCivs()

	// Game Loop
	for {
		gameplay.GenerateIntents(gameplay.GetCivs())
		// gameplay.MoveUnits(mapping.GetCivs())
		time.Sleep(time.Second * 10)
	}
}
