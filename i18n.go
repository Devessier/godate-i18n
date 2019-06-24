package i18ndate

import (
	"errors"
	"strconv"
	"time"
)

// These constants permit to specify which fields (and their string representation !) the user wants
const (
	ShortDay = 1 << iota
	LongDay
	DayNumber
	ShortMonth
	LongMonth
	MonthNumber
	Year
)

// Constants defining usual representations
const (
	BasicDate     = LongDay | DayNumber | LongMonth | Year
	NumericalDate = DayNumber | MonthNumber | Year
)

// Lang reprensents an available language
type Lang string

var availableLanguages = [...]Lang{"fr"}

// LangDay is a string representing a day in a foreign language
type LangDay string

// LangMonth is a string representing a day in a foreign language
type LangMonth string

// Translator is the structure containing the months and days slices for a given language
type Translator struct {
	Lang   Lang
	Days   []LangDay
	Months []LangMonth
}

// NewTranslator returns a translator corresponding to a given language or an error if this language is unavailable
func NewTranslator(lang string) (*Translator, error) {
	switch lang {
	case "fr":
		return NewFrenchTranslator(), nil
	default:
		return nil, errors.New("This language is unavailable :-(")
	}
}

// Returns the weekday corresponding to weekday number (Sunday is 0 for Golang Time.Weekday type)
func (tr *Translator) weekday(weekday int) string {
	var index int
	if weekday == 0 {
		index = 6
	} else {
		index = weekday - 1
	}
	return string(tr.Days[index])
}

// Returns the month corresponding to month number
func (tr *Translator) month(month int) string {
	return string(tr.Months[month-1])
}

// Translate takes a time.Time in parameter and returns a formatted string according to the provided pattern (for instance BasicDate)
func (tr *Translator) Translate(time time.Time, pattern int) (date string) {
	var weekday = time.Weekday()
	var day = time.Day()
	var month = time.Month()
	var year = time.Year()

	if (pattern & NumericalDate) != 0 {
		date = padStringNumber(day) + "/" + padStringNumber(int(month)) + "/" + strconv.Itoa(year)
		return
	}

	if (pattern&ShortDay) != 0 || (pattern&LongDay) != 0 {
		day := tr.weekday(int(weekday))
		if pattern&ShortDay != 0 {
			date += day[:3]
		} else {
			date += day
		}
	}
	if pattern&DayNumber != 0 {
		if nb := strconv.Itoa(day); len(date) > 0 {
			date += " " + nb
		} else {
			date += nb
		}
	}
	if (pattern&ShortMonth) != 0 || (pattern&LongMonth) != 0 {
		month := tr.month(int(month))
		if len(date) > 0 {
			date += " "
		}
		if pattern&ShortMonth != 0 {
			date += month[:3]
		} else {
			date += month
		}
	}
	if pattern&Year != 0 {
		if len(date) > 0 {
			date += " "
		}
		date += strconv.Itoa(year)
	}
	return
}

func padStringNumber(n int) string {
	str := strconv.Itoa(n)
	if len(str) == 1 {
		return "0" + str
	}
	return str
}
