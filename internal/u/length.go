package u

import "fmt"

func FormatLength(lengthM int) string {
	km := float64(lengthM) / 1000.0
	return fmt.Sprintf("%.1f", km)
}
