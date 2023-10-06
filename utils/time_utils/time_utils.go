package time_utils

import "time"

const defaultTimeLayout = "2006-01-02T15:04:05.999999Z"

func FormatDateTimeStringWithLayout(dataTime, layout string) (time.Time, error) {
	return time.Parse(layout, dataTime)
}

func FormatDateTimeString(timeStr string) (time.Time, error) {
	return FormatDateTimeStringWithLayout(timeStr, defaultTimeLayout)
}
