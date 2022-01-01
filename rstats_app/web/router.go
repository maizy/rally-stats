package web

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app/db"
)

const staticPrefix = "/static/"
const staticCacheMaxAge = 30 * 24 * 60 * 60

//go:embed static/*
var staticFS embed.FS

func AppendRouters(engine *gin.Engine, db *db.DBContext, devMode bool) {

	engine.GET(staticPrefix+"*filepath", StaticsHandler(devMode))

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/by-day")
	})

	engine.GET("/by-day", BuildByDayHandler(db))

	engine.GET("/version", BuildVersionHandler())
}
