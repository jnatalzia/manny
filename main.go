package main

import (
	"encoding/json"
	"strings"
	"time"

	"./gameplay"
	"./utils"
	"github.com/valyala/fasthttp"
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

	staticHandler := fasthttp.FSHandler("./webpublic", 0)

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		pth := string(ctx.Path())
		switch pth {
		default:
			splitPth := strings.Split(pth, ".")
			if len(splitPth) > 1 {
				ext := splitPth[1]
				if strings.Contains(ext, "css") {
					ctx.SetContentType("text/css")
				} else if strings.Contains(ext, "js") {
					ctx.SetContentType("application/javascript")
				} else if strings.Contains(ext, "html") {
					ctx.SetContentType("text/html")
				}
			}
			staticHandler(ctx)
		}
	}

	fasthttp.ListenAndServe(":8080", requestHandler)
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
