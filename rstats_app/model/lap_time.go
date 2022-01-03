package model

import (
	"fmt"
	"math"
	"time"

	"dev.maizy.ru/rstats/internal/u"
	"dev.maizy.ru/rstats/rstats_app/dicts"
)

type LapTime struct {
	Track     dicts.Track
	Car       dicts.Car
	StartedAt float64
	Time      float64
	TopSpeed  float64
}

func (l LapTime) TimeFormatted() string {
	return u.FormatRallyTime(l.Time)
}

func (l LapTime) StartedAsTime() time.Time {
	return time.UnixMicro(int64(l.StartedAt * math.Pow(10.0, 6.0)))
}

func (l LapTime) StartedFormatted() string {
	return l.StartedAsTime().Format("Mon Jan 2 15:04:05 MST 2006")
}

func (l LapTime) String() string {
	return fmt.Sprintf(
		"LapTime([%s] %s: %s, %s, top speed = %0.1f)",
		l.StartedFormatted(), l.TimeFormatted(), l.Track, l.Car, l.TopSpeed,
	)
}
