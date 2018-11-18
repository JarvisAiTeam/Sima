package rules

import (
 "Sima/tokens/explode"
	"Sima/tokens/stemmer"
)




type Rules struct {
	Type     string
	Punctuation map[string]string
}

func NewRules(Tokens explode.Tokens, Stemmer stemmer.AbstractStemmer, already_known bool)  *Rules{

	rules := new(Rules)

	if !already_known {
	rules.findRules(Tokens.Values, Stemmer)
	}

	return rules

}


func (rules *Rules) findRules(Tokens map[int]string, Stemmer stemmer.AbstractStemmer) {
	punctuation := GenPunctuation(Tokens, Stemmer)
	rules.Punctuation = MatchPunct(punctuation)
}


