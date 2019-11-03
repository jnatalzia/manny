package gameplay

import (
	"../utils"
)

type GamePlayer struct {
	ID        int          `json:"id"`
	Cash      int          `json:"cash"`
	Traders   []GameTrader `json:"traders"`
	TurnReady bool         `json:"ready"`
	LandColor string       `json:"land_color"`
}

type PossibleMoves struct {
	LandTilesForPurchase []int `json:"land_tiles_for_purchase"`
}

func getPlayerByID(plyrs *[]GamePlayer, id int) *GamePlayer {
	for _, p := range *plyrs {
		if p.ID == id {
			return &p
		}
	}
	utils.Logger.Panicln("No land found with requested ID")
	return &GamePlayer{}
}

// Returns adjusted player money based on actions in queue
func determinePlayerStateAfterActions(gs *GameState, pendingActions *Actions, plyr *GamePlayer) (int, []int) {
	var boughtLandIDs []int
	currentMoney := plyr.Cash
	traders := gs.Traders
	for _, act := range pendingActions.BuyCrop {
		if act.PlayerID == plyr.ID {
			cropCost := GetCostForTraderCrop(getTraderByID(&traders, act.TraderID))
			utils.Logger.Printf("Subtracting money for bought crop: %d\n", cropCost)
			currentMoney -= cropCost
		}
	}
	for _, act := range pendingActions.BuyLand {
		if act.PlayerID == plyr.ID {
			landCost := act.BidPrice
			utils.Logger.Printf("Subtracting money for bought land: %d\n", landCost)
			currentMoney -= landCost
			boughtLandIDs = append(boughtLandIDs, act.LandID)
		}
	}
	for _, act := range pendingActions.CreateTrader {
		if act.PlayerID == plyr.ID {
			utils.Logger.Printf("Subtracting money for bought trader: %d\n", traderCost)
			currentMoney -= traderCost
		}
	}
	// gs.
	utils.Logger.Printf("Total remaining money after current actions resolved: %d\n", currentMoney)
	return currentMoney, boughtLandIDs
}

func PlayerCanBuyLand(gs *GameState, pendingActions *Actions, plyr *GamePlayer, action *BuyLandAction) bool {
	availableFunds, unavailableLand := determinePlayerStateAfterActions(gs, pendingActions, plyr)
	ownedLand := []int{}

	for _, l := range gs.LandTiles {
		if l.OwnerID == plyr.ID {
			ownedLand = append(ownedLand, l.ID)
		}
	}

	return action.BidPrice <= availableFunds && !utils.IntInSlice(action.LandID, unavailableLand) && !utils.IntInSlice(action.LandID, ownedLand)
}

func PlayerCanBuyTrader(gs *GameState, pendingActions *Actions, action *CreateTraderAction) bool {
	player := getPlayerByID(&gs.Players, action.PlayerID)
	availableFunds, _ := determinePlayerStateAfterActions(gs, pendingActions, player)
	return traderCost <= availableFunds
}

func PlayerCanInfectCrop(gs *GameState, action *InfectCropAction) bool {
	lnd := getLandByID(&gs.LandTiles, action.LandID)
	return lnd.Crop >= action.Amount && action.PlayerID == lnd.OwnerID
}

func PlayerCanBuyCrop(gs *GameState, pendingActions *Actions, action *BuyCropAction) bool {
	plyr := getPlayerByID(&gs.Players, action.PlayerID)
	availableFunds, _ := determinePlayerStateAfterActions(gs, pendingActions, plyr)
	tCost := GetCostForTraderCrop(getTraderByID(&gs.Traders, action.TraderID))
	lnd := getLandByID(&gs.LandTiles, action.LandID)

	return tCost <= availableFunds && lnd.OwnerID == plyr.ID
}

func PlayerCanRouteTrader(gs *GameState, tid string, plyrid int) bool {
	t := getTraderByID(&gs.Traders, tid)
	return t.OwnerID == plyrid
}
