package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app/db"
)

func BuildByDayHandler(db *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var totalTracks int
		err := db.Times.QueryRow(`select count(*) as cnt from "laptimes"`).Scan(&totalTracks)
		if err != nil {
			returnError(c, fmt.Sprintf("unable to query times: %s", err), http.StatusInternalServerError)
			return
		}
		c.HTML(http.StatusOK, "by-day.tmpl", WithCommonVars(c, gin.H{
			"totalTracks": totalTracks,
		}))
	}
}
