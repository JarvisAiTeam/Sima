package stemmer

import (
	"github.com/liderman/rustemmer"
)


type AbstractStemmer struct{
	Stemmer  rustemmer.RuStemmer
}



func  NewStemmer(lang string) AbstractStemmer {

	AS := *new(AbstractStemmer)

	switch lang {
	case "ru": AS.Stemmer = *rustemmer.New()
	}

	return AS
}
