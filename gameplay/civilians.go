package gameplay

import (
	"math"

	"../mapping"
	"../types"
	"../utils"
)

type Civilian struct {
	ID           string
	IsAI         bool
	Location     string
	MoveIntent   string
	Path         []string
	PathPosition int
}

func GeneratePathing(civs []*Civilian) {
	for _, civ := range civs {
		// Generate intent movement using pathing
		civ.Path = civ.DetermineCivPath(&mapping.GameMap)
	}
	logger.Print("Intent Generated")
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
		if neighbor == civ.MoveIntent {
			possiblePaths = append(
				possiblePaths,
				[]string{neighbor},
			)
		} else {
			pthOpts := getPathToNode(neighbor, civ.MoveIntent, gameMap, []string{neighbor}, []string{neighbor})

			possiblePaths = append(
				possiblePaths,
				pthOpts...,
			)
		}
	}

	var shortestPath []string
	shortestPathLen := math.MaxInt64
	for _, p := range possiblePaths {
		pthLen := len(p)
		if pthLen < shortestPathLen {
			shortestPathLen = pthLen
			shortestPath = p
			continue
		}

		if pthLen == shortestPathLen && utils.GetRandomInt(100) <= 50 {
			shortestPath = p
		}
	}

	logger.Printf("Chose path %v for going from %s to %s\n", shortestPath, civ.Location, civ.MoveIntent)

	return shortestPath
}
