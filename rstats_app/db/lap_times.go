package db

import (
	"fmt"

	"dev.maizy.ru/rstats/rstats_app/model"
)

func GetAllStageTimes(db *DBContext) ([]model.StageTime, error) {
	var stageTimes []model.StageTime
	rows, err := db.Times.Query("select Track, Car, Timestamp, Time, Topspeed from laptimes order by Timestamp desc")
	if err != nil {
		return nil, fmt.Errorf("unable to get stage times: %w", err)
	}
	var i int
	for rows.Next() {
		i++
		var trackId, carId int
		var timestamp, time, topSpeed float64
		if err := rows.Scan(&trackId, &carId, &timestamp, &time, &topSpeed); err != nil {
			return nil, fmt.Errorf("unable to parse stage time #%d: %w", i, err)
		}
		stageTimes = append(stageTimes, model.StageTime{
			Track:     db.Dicts.GetTrackById(trackId),
			Car:       db.Dicts.GetCarById(carId),
			StartedAt: timestamp,
			Time:      time,
			TopSpeed:  topSpeed,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate on stage times results: %w", err)
	}
	return stageTimes, nil
}
