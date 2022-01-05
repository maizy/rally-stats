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
			StageTimes  []model.StageTime
			TotalTime   float64
			TotalLength int
		}

		stageTimes, err := db.GetAllStageTimes(dbCtx)
		if err != nil {
			returnError(c, fmt.Sprintf("unable to query stage times: %s", err), http.StatusInternalServerError)
			return
		}

		var byDays []ByDay
		for _, stageTime := range stageTimes {
			day := u.TimeToDate(stageTime.StartedAtAsTime())
			if len(byDays) == 0 || byDays[len(byDays)-1].Day != day {
				byDays = append(byDays, ByDay{
					Day:        day,
					StageTimes: []model.StageTime{stageTime},
				})
			} else {
				byDays[len(byDays)-1].StageTimes = append(byDays[len(byDays)-1].StageTimes, stageTime)
			}
			byDays[len(byDays)-1].TotalTime += stageTime.Time
			byDays[len(byDays)-1].TotalLength += stageTime.Track.Length
		}

		c.HTML(http.StatusOK, "by-day.tmpl", WithCommonVars(c, gin.H{
			"totalStagetimes": len(stageTimes),
			"byDays":          byDays,
		}))
	}
}
