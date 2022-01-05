package u

import (
	"fmt"
	"math"
	"time"
)

func FormatRallyTime(time float64) string {
	negative := ""
	if time < 0.0 {
		time = math.Abs(time)
		negative = "-"
	}
	totalMs := int(math.Round(time * 1000.0))
	ms := totalMs % 1000
	sec := (totalMs / 1000) % 60
	min := totalMs / 60000
	var hours string
	if min >= 60 {
		hours = fmt.Sprintf("%d:", min/60)
		min = min % 60
	}
	return fmt.Sprintf("%s%s%02d:%02d.%03d", negative, hours, min, sec, ms)
}

func TimeToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
