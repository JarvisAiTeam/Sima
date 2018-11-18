package rules

import (
	"Sima/tokens/stemmer"
	"github.com/fiam/gounidecode/unidecode"
	"sort"
)


var PunctuationMarks = []string{"comma","dot"}



type Punctuation struct {
	Frequency float64
	Distance  float64
	Score     int64
	Key       string
	Value     string
}

type Punctuations []Punctuation


func GenPunctuation(Tokens map[int]string, S stemmer.AbstractStemmer) Punctuations{

	punctuations := make(Punctuations,0,len(Tokens))

	frequency_counter := make(map[string]int)

	tokensLen := len(Tokens)

	i := 0

	for _,token := range Tokens {
		punct := new(Punctuation)
		punct.Value = getLastSymbol(unidecode.Unidecode(token))


		if S.Stemmer.GetWordBase(token) == token && punct.Value != "" {
			frequency_counter[punct.Value]++
			} else {
				continue
		}

		punct.Frequency = float64(frequency_counter[punct.Value])/float64(tokensLen)*100

		index := punctuations.findIndexByValue(punct.Value)

		punctuations = append(punctuations,*punct)

		if index == -1 {
			index = i
			i++
		} else {
		punctuations[index] = punctuations[len(punctuations)-1] // Copy last element to index i
		punctuations[len(punctuations)-1] = Punctuation{}  // Erase last element (write zero value)
		punctuations = punctuations[:len(punctuations)-1]

		}
	}

return punctuations.sort()
}



func MatchPunct(punctuations Punctuations) map[string]string {

	ReadyPunct := make(map[string]string)


	for i,mark := range PunctuationMarks {
        ReadyPunct[mark] = punctuations[i].Value
		}
		return ReadyPunct
}






func getLastSymbol(token string) string{
	runes := []rune(token)
	if len(runes) > 0 {return string(runes[len(runes)-1:])} else { return ""}
}



func (P Punctuations) sort() Punctuations{
	Desc(P)
	return P
}



func (P *Punctuations) findIndexByValue(value string) int{
	for i,punct := range *P {
		if punct.Value == value {return i; break}
	}
	return -1
}




func Desc(P Punctuations) {
	sort.Slice(P[:], func(i, j int) bool {
		return P[i].Frequency > P[j].Frequency

		//return P[i].FirstName < P[j].FirstName
	})
}