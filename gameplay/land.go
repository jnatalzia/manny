package gameplay

import "../utils"

type LandTile struct {
	ID           int `json:"id"`
	OwnerID      int `json:"owner_id"`
	Crop         int `json:"crop"`
	InfectedCrop int `json:"infected_crop"`
}

var (
	cropCost         = 4
	infectedCropCost = 1
)

func getLandByID(tiles *[]LandTile, id int) *LandTile {
	for _, lt := range *tiles {
		if lt.ID == id {
			return &lt
		}
	}
	utils.Logger.Panicln("No land found with requested ID")
	return &LandTile{}
}

func getBaseCostForLand(land *LandTile) int {
	return (land.Crop * cropCost) + (land.InfectedCrop * infectedCropCost)
}
