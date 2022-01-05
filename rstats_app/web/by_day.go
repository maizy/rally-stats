package web

import (
	"fmt"
	"net/http"
	"time"

	"dev.maizy.ru/rstats/internal/u"
	"dev.maizy.ru/rstats/rstats_app/model"
	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app/db"
)

func BuildByDayHandler(dbCtx *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		type ByDay struct {
			Day         time.Time
			LapTimes    []model.LapTime
			TotalTime   float64
			TotalLength int
		}

		lapTimes, err := db.GetAllLapTimes(dbCtx)
		if err != nil {
			returnError(c, fmt.Sprintf("unable to query lap times: %s", err), http.StatusInternalServerError)
			return
		}

		var byDays []ByDay
		for _, lapTime := range lapTimes {
			day := u.TimeToDate(lapTime.StartedAtAsTime())
			if len(byDays) == 0 || byDays[len(byDays)-1].Day != day {
				byDays = append(byDays, ByDay{
					Day:      day,
					LapTimes: []model.LapTime{lapTime},
				})
			} else {
				byDays[len(byDays)-1].LapTimes = append(byDays[len(byDays)-1].LapTimes, lapTime)
			}
			byDays[len(byDays)-1].TotalTime += lapTime.Time
			byDays[len(byDays)-1].TotalLength += lapTime.Track.Length
		}

		c.HTML(http.StatusOK, "by-day.tmpl", WithCommonVars(c, gin.H{
			"totalLaptimes": len(lapTimes),
			"byDays":        byDays,
		}))
	}
}
