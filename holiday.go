package goHolidays

import (
	"time"
)

type Holiday interface {
	Name() string
	Describe() string
	Check(target time.Time) bool
}

type FixedHoliday struct {
	NameStr     string
	Description string
	Month       time.Month
	Day         int
}

func (fh FixedHoliday) Name() string {
	return fh.NameStr
}

func (fh FixedHoliday) Describe() string {
	return fh.Description
}

func (fh FixedHoliday) Check(target time.Time) bool {
	if target.Month() != fh.Month {
		return false
	}

	if target.Day() != fh.Day {
		return false
	}

	return true
}

type FloatingHoliday struct {
	NameStr     string
	Description string
	Ordinal     int
	DayOfWeek   time.Weekday
	Month       time.Month
}

func (fh FloatingHoliday) Name() string {
	return fh.NameStr
}

func (fh FloatingHoliday) Describe() string {
	return fh.Description
}

func (fh FloatingHoliday) Check(target time.Time) bool {
	var minDay, maxDay int

	if target.Month() != fh.Month {
		return false
	}

	if target.Weekday() != fh.DayOfWeek {
		return false
	}

	if fh.Ordinal == -1 {
		maxDay = time.Date(target.Year(), fh.Month+1, 0, 0, 0, 0, 0, time.UTC).Day() + 1
		minDay = maxDay - 8
	} else {
		minDay = (fh.Ordinal - 1) * 7
		maxDay = minDay + 8
	}

	return target.Day() > minDay && target.Day() < maxDay
}

// Predefined Holidays

var NewYearsDay = FixedHoliday{
	"New Years Day",
	"January 1st",
	time.Month(1),
	1,
}

var MLKHoliday = FloatingHoliday{
	"Martin Luther King Holiday",
	"Second Monday in January",
	3,
	time.Monday,
	time.January,
}

var PresidentsDay = FloatingHoliday{
	"Presidents Day",
	"Second Monday in February",
	3,
	time.Monday,
	time.February,
}

var MemorialDay = FloatingHoliday{
	"Memorial Day",
	"Last Monday in May",
	-1,
	time.Monday,
	time.May,
}

var Juneteenth = FixedHoliday{
	"Juneteenth",
	"June 19th",
	time.June,
	19,
}

var IndependenceDay = FixedHoliday{
	"Independence Day",
	"July 4th",
	time.July,
	4,
}

var LaborDay = FloatingHoliday{
	"Labor Day",
	"First Monday in September",
	1,
	time.Monday,
	time.September,
}

var ColumbusDay = FloatingHoliday{
	"Columbus Day / Indigenous People Day",
	"Second Monday in October",
	2,
	time.Monday,
	time.October,
}

var VeteransDay = FixedHoliday{
	"Veterans Day",
	"11th of November",
	time.November,
	11,
}

var Thanksgiving = FloatingHoliday{
	"Thanksgiving",
	"4th thursday in November",
	4,
	time.Thursday,
	time.November,
}

var Christmas = FixedHoliday{
	"Christmas",
	"December 25th",
	time.December,
	25,
}

// Predefined HolidayList

// var UsFederalHolidays = HolidayList{[]Holiday{
// 	NewYearsDay,
// 	MLKHoliday,
// 	PresidentsDay,
// 	MemorialDay,
// 	Juneteenth,
// 	IndependenceDay,
// 	LaborDay,
// 	ColumbusDay,
// 	VeteransDay,
// 	Thanksgiving,
// 	Christmas,
// }}
