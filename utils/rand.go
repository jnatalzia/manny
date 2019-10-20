package utils

import (
	"math/rand"
	"time"

	"../types"
)

func InitRand() {
	rand.Seed(time.Now().Unix())
}

func GetRandomLocID(arr []types.UnitLocationID) types.UnitLocationID {
	return arr[rand.Intn(len(arr))]
}
