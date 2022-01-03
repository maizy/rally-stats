package u

import (
	"fmt"
	"math"
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
	min := (totalMs / 60000) % 1000
	return fmt.Sprintf("%s%02d:%02d.%03d", negative, min, sec, ms)
}
