package goHolidays

import (
	"time"
)

type Holidays interface {
	Name() string
	Describe() string
	Check(target time.Time) bool
}

type HolidayList struct {
	NameStr     string
	Description string
	Items       []Holiday
}

func (hl HolidayList) Name() string {
	return hl.NameStr
}

func (hl HolidayList) Check(dateVal time.Time) bool {
	for _, holiday := range hl.Items {
		if holiday.Check(dateVal) {
			return true
		}
	}

	return false
}

func (hl HolidayList) Describe() string {
	return hl.Description
}

// Predefined Holidays

var USFederalHolidays = HolidayList{
	"US Federal Holidays",
	"US Federal Holidays",
	[]Holiday{
		NewYearsDay,
		MLKHoliday,
		PresidentsDay,
		MemorialDay,
		Juneteenth,
		IndependenceDay,
		LaborDay,
		ColumbusDay,
		VeteransDay,
		Thanksgiving,
		Christmas,
	},
}
