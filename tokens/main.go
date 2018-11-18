package tokens

import (
 "Sima/tokens/explode"
	"Sima/tokens/rules"
	_ "Sima/recognizer"
	"Sima/tokens/stemmer"
	"Sima/tokens/sentence"
	"Sima/structs"
	"strings"
	"sync"
)




type Text struct {
	Input     string
	Lang      string
	Tokens    explode.Tokens
	Rules     rules.Rules
	Stemmer   stemmer.AbstractStemmer
	Sentences sentence.Sentences
	Emotion   string
	Type      string

}


func NewText(Input string, TModel *structs.TextModel) *Text{
	text := new(Text)
	text.InputText(Input)
	text.Lang = text.getLang(Input)
	text.Stemmer = stemmer.NewStemmer(text.Lang)
	text.Tokens = *explode.NewTokens(text.Input)


	if TModel.IsEmpty() {
		text.initRules()
	} else {
		text.setRules(TModel.Rules)
	}

	text.toSentences()

	return text
}


func (text *Text) InputText(Input string) {
	text.Input = Input
}

func (text *Text) GetText() string{
	return text.Input
}


func (text *Text) getLang(Input string) string{
	return "ru"//recognizer.GetLang(Input)
}


func (text *Text) toTokens() explode.Tokens{
   return text.Tokens
}


func (text *Text) initRules() {
	text.Rules = *rules.NewRules(text.Tokens, text.Stemmer,  false)
}

func (text *Text) setRules(Rules rules.Rules) {
	text.Rules = Rules
	}

func (text *Text) toSentences() {
	var wg sync.WaitGroup

	sent_arr := strings.Split(text.Input, text.Rules.Punctuation["dot"])

	wg.Add(len(sent_arr))

	sentences := make(sentence.Sentences,len(sent_arr))

	for i,s := range sent_arr  {
		sentences[i] = *sentence.NewSentence(s)
		sentences[i].Prepare(text.Rules.Punctuation)
		go sentences[i].ToSpeechParts()
	}
	text.Sentences = sentences
}
