package model

import (
	"fmt"
	"math"
	"time"

	"dev.maizy.ru/rstats/internal/u"
	"dev.maizy.ru/rstats/rstats_app/dicts"
)

type StageTime struct {
	Track     dicts.Track
	Car       dicts.Car
	StartedAt float64
	Time      float64
	TopSpeed  float64
}

func (l StageTime) TimeFormatted() string {
	return u.FormatRallyTime(l.Time)
}

func (l StageTime) StartedAtAsTime() time.Time {
	return time.UnixMicro(int64(l.StartedAt * math.Pow(10.0, 6.0)))
}

func (l StageTime) StartedAtFormatted() string {
	return l.StartedAtAsTime().Format("Mon Jan 2 15:04:05 MST 2006")
}

func (l StageTime) StartedAtTimeFormatted() string {
	return l.StartedAtAsTime().Format("15:04")
}

func (l StageTime) String() string {
	return fmt.Sprintf(
		"StageTime([%s] %s: %s, %s, top speed = %0.1f)",
		l.StartedAtFormatted(), l.TimeFormatted(), l.Track, l.Car, l.TopSpeed,
	)
}
