package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app/db"
)

func BuildByTrackIndexHandler(dbCtx *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "by_track_index.tmpl", WithCommonVars(c, gin.H{}))
	}
}

func BuildByTrackHandler(dbCtx *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "by_track.tmpl", WithCommonVars(c, gin.H{}))
	}
}
