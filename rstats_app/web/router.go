package web

import (
	"embed"

	"github.com/gin-gonic/gin"
)

const staticPrefix = "/static/"
const staticCacheMaxAge = 30 * 24 * 60 * 60

//go:embed static/*
var staticFS embed.FS

func AppendRouters(engine *gin.Engine, devMode bool) {

	engine.GET(staticPrefix+"*filepath", StaticsHandler(devMode))

	engine.GET("/", IndexHandler())

	engine.GET("/version", BuildVersionHandler())
}
