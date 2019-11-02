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
	rows, err := db.Query("SELECT game_state FROM games LIMIT 1;")

	if err != nil {
		utils.Logger.Panicf("There was an issue connecting to the DB %v\n", err)
	}

	for rows.Next() {
		var jsonStr string
		rows.Scan(&jsonStr)

		bytes := []byte(jsonStr)
		var gs gameplay.GameState
		json.Unmarshal(bytes, &gs)

		for _, ply := range gs.Players {
			fmt.Printf("Player %d ready state %t\n", ply.ID, ply.TurnReady)
		}
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
