package web

import (
	"math/rand"
	"net/http"

	"dev.maizy.ru/rstats/rstats_app/dicts"
	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app/db"
)

func BuildByTrackIndexHandler(dbCtx *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		type CarClassWithStats struct {
			CarClass        dicts.CarClass
			TotalStageTimes int
			BestStageTime   float64
		}

		type TracksAndCarClasses struct {
			Track      dicts.Track
			CarClasses []CarClassWithStats
		}

		type TracksByLocation struct {
			Location            dicts.Location
			TracksAndCarClasses []TracksAndCarClasses
		}
		tracksByLocation := make([]TracksByLocation, 0, len(dbCtx.Dicts.TracksByLocation)+1)
		for _, location := range dbCtx.Dicts.Locations {
			tracks := dbCtx.Dicts.TracksByLocation[location]
			tracksAndCarClasses := make([]TracksAndCarClasses, 0, len(tracks))
			for _, track := range tracks {

				// FIXME: temp
				var carClasses []CarClassWithStats
				if rand.Intn(2) == 0 {
					carClasses = []CarClassWithStats{
						{dbCtx.Dicts.GetCarClassById(100), rand.Intn(50), float64(rand.Intn(120000)) / 1000.0},
						{dbCtx.Dicts.GetCarClassById(200), rand.Intn(50), float64(rand.Intn(120000)) / 1000.0},
					}
				}
				// -

				tracksAndCarClasses = append(tracksAndCarClasses, TracksAndCarClasses{track, carClasses})
			}
			tracksByLocation = append(tracksByLocation, TracksByLocation{
				location,
				tracksAndCarClasses,
			})
		}

		c.HTML(http.StatusOK, "by_track_index.tmpl", WithCommonVars(c, gin.H{
			"tracksByLocation": tracksByLocation,
		}))
	}
}

func BuildByTrackHandler(dbCtx *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "by_track.tmpl", WithCommonVars(c, gin.H{}))
	}
}
