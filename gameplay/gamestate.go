package gameplay

type GameState struct {
	ID         int          `json:"id"`
	TurnNumber int          `json:"turn_number"`
	LandTiles  []LandTile   `json:"land"`
	Players    []GamePlayer `json:"players"`
	Traders    []GameTrader `json:"traders"`
}
