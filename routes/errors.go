package routes

import "github.com/valyala/fasthttp"

func SetErrorResponse(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(400)
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("{\"status\": \"Bad request\"}"))
}
