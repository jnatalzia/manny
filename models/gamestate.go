package models

import (
	"encoding/json"

	"../gameplay"
	"../utils"
)

func GetGamestate(gameID int) gameplay.GameState {
	row := utils.DBConn.QueryRow("SELECT game_state FROM games WHERE id = $1", gameID)
	var gsStr string
	row.Scan(&gsStr)

	bytes := []byte(gsStr)
	var gs gameplay.GameState
	json.Unmarshal(bytes, &gs)
	return gs
}
