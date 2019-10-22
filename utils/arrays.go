package utils

import (
	"../types"
)

func LocationInSlide(a types.UnitLocationID, list []types.UnitLocationID) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
