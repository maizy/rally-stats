package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app/db"
)

func BuildByDayHandler(dbCtx *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		lapTimes, err := db.GetAllLapTimes(dbCtx)
		if err != nil {
			returnError(c, fmt.Sprintf("unable to query lap times: %s", err), http.StatusInternalServerError)
			return
		}
		c.HTML(http.StatusOK, "by-day.tmpl", WithCommonVars(c, gin.H{
			"totalTracks": len(lapTimes),
			"lapTimes":    lapTimes,
		}))
	}
}
