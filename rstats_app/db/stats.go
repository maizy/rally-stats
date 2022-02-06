package db

import (
	"fmt"

	"dev.maizy.ru/rstats/rstats_app/dicts"
)

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

func GetStatsByLocationTrackAndCarClass(dbCtx *DBContext) (map[dicts.Location]TracksByLocation, error) {

	// floor function available only in sqlite math extension,
	// but cast to int effictively do the same
	rows, err := dbCtx.Times.Query(
		`select 
       		cast(Car / 100 as int) * 100 as car_class_id, 
       		Track as track_id, 
       		count(1) as times, 
       		min(Time) as best_time
		 from laptimes
		 group by car_class_id, track_id
		 order by track_id, car_class_id`)
	if err != nil {
		return nil, fmt.Errorf("unable to get tracks stat by car class: %w", err)
	}

	dbResultsByTracks := make(map[dicts.Track][]CarClassWithStats)
	var i int
	for rows.Next() {
		i++
		var trackId, carClassId, times int
		var bestTime float64
		if err := rows.Scan(&carClassId, &trackId, &times, &bestTime); err != nil {
			return nil, fmt.Errorf("unable to parse stat be tracks and carclass #%d: %w", i, err)
		}
		track := dbCtx.Dicts.GetTrackById(trackId)
		carClass := dbCtx.Dicts.GetCarClassById(carClassId)
		dbResultsByTracks[track] = append(dbResultsByTracks[track], CarClassWithStats{carClass, times, bestTime})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate on tracks stat by car class results: %w", err)
	}

	var result = make(map[dicts.Location]TracksByLocation, len(dbCtx.Dicts.Locations))
	for _, location := range dbCtx.Dicts.Locations {
		var byTracks []TracksAndCarClasses
		for _, track := range dbCtx.Dicts.TracksByLocation[location] {
			trackResults := dbResultsByTracks[track]
			byTracks = append(byTracks, TracksAndCarClasses{track, trackResults})
		}
		result[location] = TracksByLocation{location, byTracks}
	}
	unknownTracksStats := dbResultsByTracks[dicts.UnknownTrack]
	if unknownTracksStats != nil {
		result[dicts.UnknownLocation] = TracksByLocation{
			dicts.UnknownLocation,
			[]TracksAndCarClasses{{dicts.UnknownTrack, unknownTracksStats}}}
	}
	return result, nil
}
