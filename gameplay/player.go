package gameplay

import (
	"fmt"

	"../utils"
)

type GamePlayer struct {
	ID        int          `json:"id"`
	Cash      int          `json:"cash"`
	Traders   []GameTrader `json:"traders"`
	TurnReady bool         `json:"ready"`
}

// Returns adjusted player money based on actions in queue
func determinePlayerMoneyAfterActions(gs *GameState, pendingActions *Actions, plyr *GamePlayer) int {
	currentMoney := plyr.Cash
	traders := gs.Traders
	for _, act := range pendingActions.BuyCrop {
		if act.PlayerID == plyr.ID {
			cropCost := getCostForTraderCrop(getTraderByID(&traders, act.TraderID))
			utils.Logger.Printf("Subtracting money for bought crop: %d\n", cropCost)
			currentMoney -= cropCost
		}
	}
	for _, act := range pendingActions.BuyLand {
		if act.PlayerID == plyr.ID {
			landCost := getCostForLand(getLandByID(&gs.LandTiles, act.LandID))
			utils.Logger.Printf("Subtracting money for bought land: %d\n", landCost)
			currentMoney -= landCost
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
	return currentMoney
}

func DeterminePossiblePlayerActions(gs *GameState, pendingActions *Actions, pid int) *Actions {
	var player GamePlayer
	for _, plyr := range gs.Players {
		if plyr.ID == pid {
			player = plyr
			break
		}
	}
	availableFunds := determinePlayerMoneyAfterActions(gs, pendingActions, &player)
	fmt.Printf("Available funds for player after action: %d\n", availableFunds)
	return &Actions{}
}
