package main

import (
	"strings"
	"time"

	"./gameplay"
	"./routes"
	"./utils"
	"github.com/valyala/fasthttp"
)

// NOTE: If we are storing in json and updating one at a time,
// we need a queue per game to host move order
// To ensure we do not overwrite existing JSON
func main() {
	utils.ConnectToDB()
	utils.InitLogger()
	utils.InitRand()

	staticHandler := fasthttp.FSHandler("./webpublic", 0)

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		pth := string(ctx.Path())
		switch pth {
		case "/gamestate":
			routes.GamestateForPlayer(ctx)
		default:
			splitPth := strings.Split(pth, ".")
			if len(splitPth) > 1 {
				ext := splitPth[1]
				if strings.Contains(ext, "css") {
					ctx.SetContentType("text/css")
				} else if strings.Contains(ext, "js") {
					ctx.SetContentType("application/javascript")
				} else if strings.Contains(ext, "html") {
					ctx.SetContentType("text/html")
				}
			}
			staticHandler(ctx)
		}
	}

	fasthttp.ListenAndServe(":8080", requestHandler)
}

func startGameLoop() {
	gameplay.BuildCivs()
	gameplay.GeneratePathing(gameplay.GetCivs())

	// Game Loop
	for {
		gameplay.TakeTurn()
		time.Sleep(time.Second * 10)
	}
}
