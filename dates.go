package easypostdevtools

import "time"

type Dates struct {
}

func (d Dates) contains(list []int, item int) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func (d Dates) IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (d Dates) GetLastDayOfMonth(month int, year int) int {
	if d.contains([]int{1, 3, 5, 7, 8, 10, 12}, month) {
		return 31
	} else if d.contains([]int{4, 6, 9, 11}, month) {
		return 30
	} else {
		if d.IsLeapYear(year) {
			return 29
		} else {
			return 28
		}
	}
}

func (d Dates) IsLastDayOfMonth(date time.Time) bool {
	return date.Day() == d.GetLastDayOfMonth(int(date.Month()), date.Year())
}

func (d Dates) IsLastMonthOfYear(date time.Time) bool {
	return date.Month() == time.December
}

func (d Dates) IsLastDayOfYear(date time.Time) bool {
	return date.Month() == time.December && date.Day() == 31
}

func (d Dates) GetFutureDateThisYear() time.Time {
	if d.IsLastDayOfYear(time.Now()) {
		panic("This year is over.")
	}
	month := 0
	if d.IsLastDayOfMonth(time.Now()) {
		// pull from next month on
		month = Random{}.GetRandomIntInRange(int(time.Now().Month())+1, 12)
	} else {
		// pull from next day on
		month = Random{}.GetRandomIntInRange(int(time.Now().Month()), 12)
	}

	day := 0
	maxDays := d.GetLastDayOfMonth(month, time.Now().Year())
	if month == int(time.Now().Month()) {
		// pull from tomorrow on
		day = Random{}.GetRandomIntInRange(int(time.Now().Day())+1, maxDays)
	} else {
		// pull from day 1 on
		day = Random{}.GetRandomIntInRange(1, maxDays)
	}

	return time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (d Dates) GetFutureDateThisMonth() time.Time {
	if d.IsLastDayOfMonth(time.Now()) {
		panic("This month is over.")
	}
	day := Random{}.GetRandomIntInRange(int(time.Now().Day())+1, d.GetLastDayOfMonth(int(time.Now().Month()), time.Now().Year()))
	return time.Date(time.Now().Year(), time.Now().Month(), day, 0, 0, 0, 0, time.UTC)
}

func (d Dates) GetDateAfter(date time.Time) time.Time {
	return date.AddDate(0, 0, Random{}.GetRandomIntInRange(1, 30))
}

func (d Dates) GetDateBefore(date time.Time) time.Time {
	return date.AddDate(0, 0, -Random{}.GetRandomIntInRange(1, 30))
}

func (d Dates) GetFutureDates(count int) []time.Time {
	var dates []time.Time
	currentDate := time.Now()
	for i := 0; i < count; i++ {
		currentDate = d.GetDateAfter(currentDate)
		dates = append(dates, currentDate)
	}
	return dates
}

func (d Dates) GetPastDates(count int) []time.Time {
	var dates []time.Time
	currentDate := time.Now()
	for i := 0; i < count; i++ {
		currentDate = d.GetDateBefore(currentDate)
		dates = append(dates, currentDate)
	}
	return dates
}
