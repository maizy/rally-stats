package web

import (
	"log"
	"net/http"

	"dev.maizy.ru/rstats/rstats_app/dicts"
	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app/db"
)

func BuildByTrackIndexHandler(dbCtx *db.DBContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var tracksByLocation []db.TracksByLocation
		tracksByLocationStat, err := db.GetStatsByLocationTrackAndCarClass(dbCtx)
		if err != nil {
			log.Printf("Unable to fetch tracks stats: %s", err)
			returnError(c, "Unable to fetch tracks stats", http.StatusInternalServerError)
			return
		}
		// use locations with predefined order
		for _, location := range dbCtx.Dicts.Locations {
			if tracksAndCarClasses, exists := tracksByLocationStat[location]; exists {
				tracksByLocation = append(tracksByLocation, tracksAndCarClasses)
			}
		}

		if unknownLocationStats, exists := tracksByLocationStat[dicts.UnknownLocation]; exists {
			tracksByLocation = append(tracksByLocation, unknownLocationStats)
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
