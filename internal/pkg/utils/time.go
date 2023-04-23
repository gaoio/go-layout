package utils

import "time"

// LocCST 东八区
func LocCST() *time.Location {
	return time.FixedZone("CST", 8*3600)
}

const (
	TimeLayout = "2006-01-02 15:04:05"
)

func TimeZeroStringFormat(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Add(8 * time.Hour).Format(TimeLayout)
}
