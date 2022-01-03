package u

import "testing"

func TestFormatRallyTime(t *testing.T) {
	tests := []struct {
		name string
		time float64
		want string
	}{
		{"simple", 90.5, "01:30.500"},
		{"zero", 0.0, "00:00.000"},
		{"zero ms", 90.0, "01:30.000"},
		{"only seconds", 43.123, "00:43.123"},
		{"leading zeroes", 5.001, "00:05.001"},
		{"too precise ms", 124.0009, "02:04.001"},
		{"one ms before", 59.999, "00:59.999"},
		{"one ms/10 before", 59.9999, "01:00.000"},
		{"more than 99 min", 6121.1, "102:01.100"},
		{"negative", -90.5, "-01:30.500"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatRallyTime(tt.time); got != tt.want {
				t.Errorf("FormatRallyTime(%f) = %v, want %v", tt.time, got, tt.want)
			}
		})
	}
}
