package util

import (
	"time"
)

// ISOWeekRange ...
func ISOWeekRange(year int, week int) (time.Time, time.Time) {
	PST, _ := time.LoadLocation("America/Los_Angeles")
	date := time.Date(year, 0, 0, 0, 0, 0, 0, PST)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date, date.AddDate(0,0,7).Add(-1*time.Second)
}