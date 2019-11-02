package gameplay

import "../utils"

type GameTrader struct {
	ID            string `json:"id"`
	OwnerID       int    `json:"owner_id"`
	LocationID    int    `json:"location_id"`
	SafeCrop      int    `json:"safe_crop"`
	InfectedCrop  int    `json:"infected_crop"`
	DestinationID int    `json:"destination_id"`
}

var (
	traderCost     = 50
	traderCropCost = 4
)

func getTraderByID(traders *[]GameTrader, id string) *GameTrader {
	utils.Logger.Printf("Finding trader with id: %s", id)
	for _, t := range *traders {
		if t.ID == id {
			return &t
		}
	}
	utils.Logger.Panicln("No trader found with requested ID")
	return &GameTrader{}
}

func getCostForTraderCrop(trader *GameTrader) int {
	return (trader.SafeCrop + trader.InfectedCrop) * traderCropCost
}
