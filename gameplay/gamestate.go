package gameplay

type GameState struct {
	ID         int          `json:"id"`
	TurnNumber int          `json:"turn_number"`
	LandTiles  []LandTile   `json:"land"`
	Players    []GamePlayer `json:"players"`
	AI         GameAI       `json:"ai"`
}

type LandTile struct {
	ID          int `json:"id"`
	OwnerFarmID int `json:"owner_farm_id"`
}

type GamePlayer struct {
	ID           int              `json:"id"`
	FarmID       int              `json:"farm_id"`
	Cash         int              `json:"cash"`
	Traders      []GameTrader     `json:"traders"`
	OpenRequests []PlayerRequests `json:"requests"`
	TurnReady    bool             `json:"ready"`
}

type GameTrader struct {
	LocationID    int  `json:"location_id"`
	IsMalicious   bool `json:"is_malicious"`
	DestinationID int  `json:"destination_id"`
}

type PlayerRequests struct {
	RequestType   string `json:"type"`
	DestinationID int    `json:"destination_id"`
}

type GameAI struct {
	Traders []GameTrader `json:"traders"`
}
