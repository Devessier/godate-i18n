package i18ndate

// The days in french
var frenchDays = []LangDay{
	"Lundi",
	"Mardi",
	"Mercredi",
	"Jeudi",
	"Vendredi",
	"Samedi",
	"Dimanche",
}

// The months in french
var frenchMonths = []LangMonth{
	"Janvier",
	"Février",
	"Mars",
	"Avril",
	"Mai",
	"Juin",
	"Juillet",
	"Août",
	"Septembre",
	"Novembre",
	"Décembre",
}

// NewFrenchTranslator returns a *Translator with the french days and months
func NewFrenchTranslator() *Translator {
	return &Translator{
		Lang:   "fr",
		Days:   frenchDays,
		Months: frenchMonths,
	}
}
