package train

import (
	"fmt"
	"bufio"
	"Sima/tokens"
	"Sima/tokens/rules"
	"Sima/database"
	"encoding/json"
	"os"
	"Sima/structs"
)





type Marks struct {
	Value string
	Counter int
	Rule string
}

type Res []Marks




func TrainTokenization(scanner *bufio.Scanner) {

	//myscanner := bufio.NewScanner(os.Stdin) //os.StdIn

	myscanner := scanner


	var i,d,c float64

	Res := make([]Marks,0,5)

	forSorts := make(map[string]int)

	Result := make(map[string]string)

	var Lang string

	for   {
		S1 := ScanInput(myscanner)

		if S1 == "" {break}

		if S1 == "-" {continue}

        i++


		Text := tokens.NewText(S1,&structs.TextModel{})

		Lang = Text.Lang

		Rules := Text.Rules

		for _,rule := range rules.PunctuationMarks {

			index := inArray(Res, Rules.Punctuation[rule])

			mark := new(Marks)
			mark.Value = Rules.Punctuation[rule]

			if index == -1 {
				mark.Rule = rule
				mark.Counter = 1
				Res = append(Res, *mark)
			} else {
				mark.Rule = Res[index].Rule
				mark.Counter = Res[index].Counter+1
				Res[index] = *mark
			}

			if forSorts[mark.Rule] < mark.Counter {forSorts[mark.Rule] = mark.Counter}
		}



		if Text.Rules.Punctuation["comma"] == "," {c++}
		if Text.Rules.Punctuation["dot"] == "." {d++}


	}

	//Result := findBiggest(Res)
	for _,rule := range rules.PunctuationMarks {
		Result[rule] = findRuleByCounter(Res, forSorts[rule],rule).Value
	}
	fmt.Printf("%f comma accuracy \n",(c/i)*100.00)
	fmt.Printf("%f dot accuracy",(d/i)*100.00)

	fmt.Println()
	fmt.Println(Result)


	name := "text_model_"+Lang

	Rules := *new(rules.Rules)
	Rules.Punctuation = Result
	insertTrainingResult(&structs.TextModel{Rules},name)
}


func TrainTextModel() {
	file,_ := os.Open("train/input.txt")
	myscanner := bufio.NewScanner(file) //os.StdIn

	TrainTokenization(myscanner)
}


func insertTrainingResult(Result *structs.TextModel, name string) {
	jsonResult,err := json.Marshal(Result)

	if err != nil {
		panic(err)
	}

	err = database.GetRedisClient().Set(name,jsonResult,0).Err()

	if err != nil {
		panic(err)
	}
}


func GetTrainingResult(lang string) *structs.TextModel{
	var TModel *structs.TextModel

	key := "text_model_"+lang
	if database.Exists(key) == 1 {
		val, err := database.GetRedisClient().Get(key).Result()
		if err != nil {
			panic(err)
		}
		TModel = new(structs.TextModel)
		json.Unmarshal([]byte(val), &TModel)
	} else {
        TrainTextModel()
		TModel = GetTrainingResult(lang)
	}
	return TModel
}


func inArray(arr []Marks,input string)  int{
	for i,mark := range arr {
		if mark.Value == input {
			return i
		}
	}
	return -1
}

func findRuleByCounter(arr []Marks, counter int, rule string)  Marks{
	for _,item := range arr {
		if item.Counter == counter && item.Rule == rule {
			return item
		}
	}
	return arr[0]
}


func ScanInput(myscanner *bufio.Scanner)  string{
	myscanner.Scan()
	return myscanner.Text()
}




















