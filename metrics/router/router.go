package router

import (
	"github.com/atompi/go-kits/metrics/handler"
	"github.com/gin-gonic/gin"
)

func MetricsRouter(routerGroup *gin.RouterGroup, metricsPath string) {
	routerGroup.GET(metricsPath, handler.NewPromHandler())
}
