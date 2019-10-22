package utils

import (
	"math/rand"
	"time"
)

func InitRand() {
	rand.Seed(time.Now().Unix())
}

func GetRandomLocID(arr []string) string {
	return arr[rand.Intn(len(arr))]
}

func GetRandomInt(topBound int) int {
	return rand.Intn(topBound)
}
