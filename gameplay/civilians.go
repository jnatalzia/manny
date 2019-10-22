package gameplay

import (
	"errors"

	"../mapping"
	"../types"
	"../utils"
)

type Civilian struct {
	IsAI       bool
	Location   types.UnitLocationID
	MoveIntent types.UnitLocationID
	Path       []types.UnitLocationID
}

func GenerateIntents(civs []*Civilian) {
	for _, civ := range civs {
		// Generate intent movement using pathing
		civ.Path = civ.DetermineCivPath(&mapping.GameMap)
	}
	logger.Print("Intents Generated")
}

func MoveUnits(civs []*Civilian) {
	for i, civ := range civs {
		// civ.Location = civ.MoveIntent
		logger.Printf("Civilian %d moved. %s", i, civ.Location)
	}
}

func getPathToNode(
	locOne types.UnitLocationID,
	locTwo types.UnitLocationID,
	gameMap *types.Map,
	path []types.UnitLocationID,
	visited []types.UnitLocationID,
) ([]types.UnitLocationID, error) {
	gm := *gameMap
	for _, nbr := range gm[locOne].Neighbors {
		path = append(path, nbr)

		logger.Printf("Current visited: %v", visited)

		if utils.LocationInSlide(nbr, visited) {
			continue
		}

		visited = append(visited, nbr)
		if nbr == locTwo {
			return path, nil
		}

		logger.Printf("Looking farther down the rabbit hole.\nCurrent Path: %v", path)
		return getPathToNode(nbr, locTwo, gameMap, path, visited)
	}

	logger.Println("Somehow got to the block below the for loop")
	return path, errors.New("No viable path for current")
}

// Based on a unit's intent, draw a path for them
func (civ *Civilian) DetermineCivPath(gameMap *types.Map) []types.UnitLocationID {
	gm := *gameMap

	possiblePaths := [][]types.UnitLocationID{}

	for _, neighbor := range gm[civ.Location].Neighbors {
		pth, err := getPathToNode(neighbor, civ.MoveIntent, gameMap, []types.UnitLocationID{neighbor}, []types.UnitLocationID{neighbor})
		if err != nil {
			logger.Printf("No viable path found from %s to %s\n", neighbor, civ.MoveIntent)
			continue
		}
		possiblePaths = append(
			possiblePaths,
			pth,
		)
	}

	logger.Printf("For path from %s to %s the following paths exist:\n", civ.Location, civ.MoveIntent)
	var shortestPath []types.UnitLocationID
	shortestPathLen := 1000000
	for _, p := range possiblePaths {
		logger.Println(p)
		pthLen := len(p)
		if pthLen < shortestPathLen {
			shortestPathLen = pthLen
			shortestPath = p
			logger.Printf("Shorter path of length %d found\n", shortestPathLen)
			continue
		}

		if pthLen == shortestPathLen {
			logger.Printf("Path found with equal %d length route\n", pthLen)
		}
	}

	return shortestPath
}
