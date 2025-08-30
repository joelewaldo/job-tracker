package http

import (
	"time"

	"github.com/joelewaldo/job-tracker/api/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		next(ctx)
		latency := time.Since(start)

		logger.Log.WithFields(logrus.Fields{
			"method":      string(ctx.Method()),
			"path":        string(ctx.Path()),
			"status":      ctx.Response.StatusCode(),
			"remote_addr": ctx.RemoteAddr().String(),
			"latency_ms":  latency.Milliseconds(),
		}).Info("request completed")
	}
}
