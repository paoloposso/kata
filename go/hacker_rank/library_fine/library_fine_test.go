package libraryfine

import (
	"testing"
)

func TestLibraryFineBeforeDueDate(t *testing.T) {
	fine := LibraryFine(1, 12, 2022, 1, 1, 2023)
	if fine != 0 {
		t.Fail()
	}
}

func TestLibraryFineSameDueDate(t *testing.T) {
	fine := LibraryFine(1, 12, 2022, 1, 1, 2023)
	if fine != 0 {
		t.Fail()
	}
}

func TestLibraryReturnNextYearTwoDaysLate(t *testing.T) {
	fine := LibraryFine(2, 1, 2023, 31, 12, 2022)
	if fine != 10000 {
		t.Fail()
	}
}

func TestLibraryFineReturnNextYearOneDay(t *testing.T) {
	fine := LibraryFine(10, 1, 2023, 31, 12, 2022)
	if fine != 10000 {
		t.Fail()
	}
}

func TestLibraryFineReturnNextYear(t *testing.T) {
	fine := LibraryFine(1, 1, 2023, 15, 12, 2022)
	if fine != 10000 {
		t.Fail()
	}
}

func TestLibraryFineReturnSameMonth(t *testing.T) {
	fine := LibraryFine(15, 12, 2023, 1, 12, 2023)
	if fine != 210 {
		t.Fail()
	}
}

func TestDifferentMonthSameYear(t *testing.T) {
	fine := LibraryFine(2, 7, 1014, 1, 1, 1014)
	if fine != 3000 {
		t.Fail()
	}
}
