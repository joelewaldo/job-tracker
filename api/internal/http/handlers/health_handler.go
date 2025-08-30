package handlers

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func HealthHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, `{"status":"ok"}`)
}
