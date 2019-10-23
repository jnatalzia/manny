package utils

import (
	"fmt"
	"log"
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

func GetGUID() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	uuid := fmt.Sprintf("%x", b[0:4])
	return uuid
}
