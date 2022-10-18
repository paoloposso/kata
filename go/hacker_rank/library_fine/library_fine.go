package libraryfine

import (
	"time"
)

func LibraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	returnDate := time.Date(int(y1), time.Month(m1), int(d1), 0, 0, 0, 0, time.UTC)
	dueDate := time.Date(int(y2), time.Month(m2), int(d2), 0, 0, 0, 0, time.UTC)

	if returnDate.Before(dueDate) || returnDate.Equal(dueDate) {
		return 0
	} else if returnDate.Year() == dueDate.Year() && returnDate.Month() == dueDate.Month() {
		daysLate := int32(returnDate.UTC().Sub(dueDate).Hours() / 24)
		return daysLate * 15
	} else if returnDate.Year() == dueDate.Year() {
		monthsLate := int32(returnDate.Month() - dueDate.Month())
		return 500 * monthsLate
	}
	return 10000
}
