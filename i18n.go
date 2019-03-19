package i18ndate

import (
	"encoding/json"
	"io"
	"os"
)

type Lang string

const pathToTranslations = "./lang/"

var AvailableLanguages = [...]Lang{"fr"}

type LangDay string
type LangMonth string

type LangConvertor struct {
	Lang   Lang
	Days   []LangDay
	Months []LangMonth
}

// NewConvertor returns a LangConvertor struct containing everything needed to translate a date to a specific language
// If the provided lang isn't supported by the library, it loads a translation json file from pathToI18nFile
// If an error occured, it returns the error and a nil pointer, otherwise, error is nil and LangConvertor is ready to be used to translate
func NewConvertor(lang, pathToI18nFile string) (error, *LangConvertor) {
	var err error
	var foundLanguage bool = false
	for i := range AvailableLanguages {
		if Lang(lang) == AvailableLanguages[i] {
			foundLanguage = true
			break
		}
	}

	var reader io.Reader
	if !foundLanguage {
		reader, err = os.Open(pathToTranslations)
		if err != nil {
			return err, nil
		}
	} else {
		reader, err = os.Open(pathToTranslations + lang)
		if err != nil {
			return err, nil
		}
	}
	var decoder *json.Decoder = json.NewDecoder(reader)
	var convertor *LangConvertor
	if err := decoder.Decode(convertor); err != nil {
		return err, nil
	}
	return nil, convertor
}
