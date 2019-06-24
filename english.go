package i18ndate

var englishDays = [...]LangDay{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thurday",
	"Friday",
	"Saturday",
	"Sunday",
}

var englishMonths = [...]LangMonth{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// NewEnglishTranslator returns a *Translator with the english days and months
func NewEnglishTranslator() *Translator {
	return &Translator{
		Lang:   "en",
		Days:   englishDays[:],
		Months: englishMonths[:],
	}
}
