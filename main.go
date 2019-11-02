package main

import (
	"time"

	"./gameplay"
	"./mapping"
	"./utils"
)

func main() {
	utils.InitLogger()
	utils.InitRand()

	mapping.InitHelperData()
	gameplay.BuildCivs()
	gameplay.GeneratePathing(gameplay.GetCivs())
	// Game Loop
	for {
		gameplay.TakeTurn()
		time.Sleep(time.Second * 10)
	}
}
