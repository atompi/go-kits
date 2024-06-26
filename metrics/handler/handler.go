// Package gin is a helper package to get a gin compatible middleware.
package handler

import (
	"context"

	"github.com/atompi/go-kits/metrics/middleware"
	"github.com/atompi/go-kits/metrics/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewPromHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Handler returns a Gin measuring middleware.
func Handler(handlerID string) gin.HandlerFunc {
	m := middleware.New(
		middleware.Config{
			Recorder: prometheus.NewRecorder(prometheus.Config{}),
		},
	)

	return func(c *gin.Context) {
		r := &reporter{c: c}
		m.Measure(handlerID, r, func() {
			c.Next()
		})
	}
}

type reporter struct {
	c *gin.Context
}

func (r *reporter) Method() string { return r.c.Request.Method }

func (r *reporter) Context() context.Context { return r.c.Request.Context() }

func (r *reporter) URLPath() string { return r.c.Request.URL.Path }

func (r *reporter) StatusCode() int { return r.c.Writer.Status() }

func (r *reporter) BytesWritten() int64 { return int64(r.c.Writer.Size()) }
