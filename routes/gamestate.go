package routes

import (
	"encoding/json"
	"strconv"

	"github.com/valyala/fasthttp"

	"../gameplay"
	"../models"
	"../utils"
)

type TraderRes struct {
	ID            string `json:"id"`
	LocationID    int    `json:"location_id"`
	Cost          int    `json:"trader_cost"`
	DestinationID int    `json:"destination_id"`
}

// Should we include player cash here?
type PlayerRes struct {
	ID        int  `json:"id"`
	TurnReady bool `json:"ready"`
}

type CurrentPlayerRes struct {
	ID        int                   `json:"id"`
	Cash      int                   `json:"cash"`
	TurnReady bool                  `json:"ready"`
	LandTiles []gameplay.LandTile   `json:"land_tiles"`
	Traders   []gameplay.GameTrader `json:"traders"`
}

type LandTileRes struct {
	ID       int `json:"id"`
	OwnerID  int `json:"owner_id"`
	BaseCost int `json:"base_cost"`
}

type GamestateResForPlayer struct {
	CurrentPlayer CurrentPlayerRes `json:"current_player"`
	Traders       []TraderRes      `json:"traders"`
	Players       []PlayerRes      `json:"players"`
	LandTiles     []LandTileRes    `json:"land_tiles"`
	TurnNumber    int              `json:"turn_number"`
	ID            int              `json:"id"`
}

func GamestateForPlayer(ctx *fasthttp.RequestCtx) {
	qArgs := ctx.QueryArgs()

	if !qArgs.Has("gid") || !qArgs.Has("pid") {
		utils.Logger.Println("Missing required arg")
		SetErrorResponse(ctx)
		return
	}

	gameID, err := strconv.Atoi(string(qArgs.Peek("gid")))
	plyrID, errTwo := strconv.Atoi(string(qArgs.Peek("pid")))

	if err != nil || errTwo != nil {
		utils.Logger.Println(err)
		utils.Logger.Println(errTwo)
		SetErrorResponse(ctx)
		return
	}

	gs := models.GetGamestate(gameID)
	utils.Logger.Println(gs)

	// Format based on player ID
	// Require:
	// List of traders (no info on malicious or not)
	// List of players
	// List of land and associated cost (should you return infected and uninfected count if not player?)
	filteredLand := []LandTileRes{}
	playerLand := []gameplay.LandTile{}
	filteredTraders := []TraderRes{}
	playerTraders := []gameplay.GameTrader{}

	for _, lnd := range gs.LandTiles {
		if plyrID == lnd.OwnerID {
			playerLand = append(playerLand, lnd)
		}
		filteredLand = append(filteredLand, LandTileRes{
			ID:       lnd.ID,
			OwnerID:  lnd.OwnerID,
			BaseCost: gameplay.GetBaseCostForLand(&lnd),
		})

	}

	for _, t := range gs.Traders {
		if plyrID == t.OwnerID {
			playerTraders = append(playerTraders, t)
		}
		filteredTraders = append(filteredTraders, TraderRes{
			ID:            t.ID,
			LocationID:    t.LocationID,
			DestinationID: t.DestinationID,
			Cost:          gameplay.GetCostForTraderCrop(&t),
		})
	}

	pRes := CurrentPlayerRes{
		LandTiles: playerLand,
		Traders:   playerTraders,
		ID:        plyrID,
	}
	otherPlyrs := []PlayerRes{}
	for _, p := range gs.Players {
		if plyrID == p.ID {
			pRes.Cash = p.Cash
			pRes.TurnReady = p.TurnReady
		} else {
			otherPlyrs = append(otherPlyrs, PlayerRes{
				ID:        p.ID,
				TurnReady: p.TurnReady,
			})
		}
	}

	jStr, err := json.Marshal(GamestateResForPlayer{
		CurrentPlayer: pRes,
		Traders:       filteredTraders,
		Players:       otherPlyrs,
		LandTiles:     filteredLand,
		TurnNumber:    gs.TurnNumber,
		ID:            gs.ID,
	})

	if err != nil || errTwo != nil {
		utils.Logger.Println(err)
		SetErrorResponse(ctx)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(jStr)
}
