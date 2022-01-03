package db

import (
	"fmt"

	"dev.maizy.ru/rstats/rstats_app/model"
)

func GetAllLapTimes(db *DBContext) ([]model.LapTime, error) {
	var lapTimes []model.LapTime
	rows, err := db.Times.Query("select Track, Car, Timestamp, Time, Topspeed from laptimes order by Timestamp desc")
	if err != nil {
		return nil, fmt.Errorf("unable to get laptimes: %w", err)
	}
	var i int
	for rows.Next() {
		i++
		var trackId, carId int
		var timestamp, time, topSpeed float64
		if err := rows.Scan(&trackId, &carId, &timestamp, &time, &topSpeed); err != nil {
			return nil, fmt.Errorf("unable to parse laptime #%d: %w", i, err)
		}
		lapTimes = append(lapTimes, model.LapTime{
			Track:     db.Dicts.GetTrackById(trackId),
			Car:       db.Dicts.GetCarById(carId),
			StartedAt: timestamp,
			Time:      time,
			TopSpeed:  topSpeed,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate on laptimes results: %w", err)
	}
	return lapTimes, nil
}
