package main

import (
	"encoding/json"
	"time"

	"./gameplay"
	"./utils"
)

// NOTE: If we are storing in json and updating one at a time,
// we need a queue per game to host move order
// To ensure we do not overwrite existing JSON
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
