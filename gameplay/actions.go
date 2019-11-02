package gameplay

// Possible Actions
// Buy crop (from trader on land), Infect crop to sell, Create trader, Route trader, Buy land
// Three currencies, cash, crop, infected crop
// Land value: Crop = 1 coin each, Infected Crop = .25 coin each
// Buy crop at 1 coin each
// Selling good crop is only way to make money but lowers land value
// Selling infected crop is a monetary gain and lowers the overall value of enemy land (what is the downside here - land loses value?)

type Actions struct {
	BuyCrop      []BuyCropAction      `json:"buy_crop"`
	CreateTrader []CreateTraderAction `json:"create_trader"`
	InfectCrop   []InfectCropAction   `json:"infect_crop"`
	BuyLand      []BuyLandAction      `json:"buy_land"`
	RouteTrader  []RouteTraderAction  `json:"route_trader"`
}

type BuyCropAction struct {
	ID       string `json:"id"`
	PlayerID int    `json:"player_id"`
	LandID   int    `json:"land_id"`
	TraderID string `json:"trader_id"`
}

type CreateTraderAction struct {
	ID           string `json:"id"`
	PlayerID     int    `json:"player_id"`
	SafeCrop     int    `json:"safe_crop"`
	InfectedCrop int    `json:"infected_crop"`
	LandID       int    `json:"land_id"`
}

type RouteTraderAction struct {
	ID            string `json:"id"`
	PlayerID      int    `json:"player_id"`
	TraderID      int    `json:"trader_id"`
	DestinationID int    `json:"destination_id"`
}

type InfectCropAction struct {
	ID       string `json:"id"`
	PlayerID int    `json:"player_id"`
	Amount   int    `json:"amount"`
	LandID   int    `json:"land_id"`
}

type BuyLandAction struct {
	ID       string `json:"id"`
	PlayerID int    `json:"player_id"`
	LandID   int    `json:"land_id"`
}
