package main

import (
	"encoding/json"
	"fmt"
	"time"

	"./gameplay"
	"./utils"
)

func main() {
	utils.ConnectToDB()
	utils.InitLogger()
	utils.InitRand()

	db := utils.DBConn
	rows, err := db.Query("SELECT game_state, pending_moves FROM games LIMIT 1;")

	if err != nil {
		utils.Logger.Panicf("There was an issue connecting to the DB %v\n", err)
	}
	for rows.Next() {
		var gsStr string
		var pendingMoveStr string
		rows.Scan(&gsStr, &pendingMoveStr)

		bytes := []byte(gsStr)
		var gs gameplay.GameState
		json.Unmarshal(bytes, &gs)

		bytes = []byte(pendingMoveStr)
		var pa gameplay.Actions
		json.Unmarshal(bytes, &pa)

		fmt.Println(pa.BuyCrop)

		possibleActions := gameplay.DeterminePossiblePlayerActions(&gs, &pa, 0)

		fmt.Println("**********")
		fmt.Println(possibleActions)
		fmt.Println("**********")
	}

	// startGameLoop()
}

func startGameLoop() {
	gameplay.BuildCivs()
	gameplay.GeneratePathing(gameplay.GetCivs())

	// Game Loop
	for {
		gameplay.TakeTurn()
		time.Sleep(time.Second * 10)
	}
}
