package structs

import (
"reflect"
	"Sima/tokens/rules"
)


type TextModel struct{
	Rules rules.Rules `json:"rules"`
}


func (TM *TextModel) IsEmpty() bool{
	return reflect.DeepEqual(TM,&TextModel{})
}







