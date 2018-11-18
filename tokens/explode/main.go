package explode

import (
	"strings"
)




var Replacer = strings.NewReplacer(",", "", "/", "", ":","", ";","", "'","", "@","", "#","", "$","", "%","", "^","", "&","", "(","", ")","",
	"№","", "-","", "_","", "+","", "=","", "\\","", "/","", "|","", "<","", ">","", "{","", "[","", "]","", "}","")



type Tokens struct {
	Values   map[int]string
	Type     string
	Length   int

}


func NewTokens(Input string) *Tokens{
	Tokens := new(Tokens)
	Tokens.parse(Input)
	return Tokens
}



func (tokens *Tokens) parse(Input string) {
	tokensArr := strings.Split(Input, " ")

	tokens.Values = make(map[int]string)
	tokens.Length = len(tokens.Values)

	for i,token := range tokensArr {tokens.Values[i] = token}
}


