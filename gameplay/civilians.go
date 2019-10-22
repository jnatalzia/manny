package gameplay

import (
	"../mapping"
	"../types"
	"../utils"
)

type Civilian struct {
	IsAI       bool
	Location   string
	MoveIntent string
	Path       []string
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
	locOne string,
	locTwo string,
	gameMap *types.Map,
	oldPath []string,
	oldVisited []string,
) [][]string {
	gm := *gameMap
	options := [][]string{}

	for _, nbr := range gm[locOne].Neighbors {
		path := make([]string, len(oldPath))
		copy(path, oldPath)
		path = append(path, nbr)

		if utils.LocationInSlice(nbr, oldVisited) {
			continue
		}

		visited := make([]string, len(oldVisited))
		copy(visited, oldVisited)
		visited = append(visited, nbr)
		if nbr == locTwo {
			options = append(options, path)
			continue
		}

		for _, opt := range getPathToNode(nbr, locTwo, gameMap, path, visited) {
			options = append(options, opt)
		}
	}

	return options
}

// Based on a unit's intent, draw a path for them
func (civ *Civilian) DetermineCivPath(gameMap *types.Map) []string {
	gm := *gameMap

	possiblePaths := [][]string{}

	for _, neighbor := range gm[civ.Location].Neighbors {
		pthOpts := getPathToNode(neighbor, civ.MoveIntent, gameMap, []string{neighbor}, []string{neighbor})

		possiblePaths = append(
			possiblePaths,
			pthOpts...,
		)
	}

	logger.Printf("For path from %s to %s the following paths exist:\n", civ.Location, civ.MoveIntent)
	var shortestPath []string
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

		if pthLen == shortestPathLen && utils.GetRandomInt(100) <= 50 {
			logger.Printf("Path found with equal %d length route\n Check passed, replacing\n", pthLen)
			shortestPath = p
		}
	}

	logger.Printf("Chose path %v\n", shortestPath)

	return shortestPath
}
