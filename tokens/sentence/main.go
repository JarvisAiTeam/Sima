/* POS that we can understand :
   NOUN - СУЩЕСТВИТЕЛЬНОЕ
   VERB - ГЛАГОЛ
   ADVERB - НАРЕЧИЕ
   PRONOUN - МЕСТОИМЕНИЕ
   ADJECTIVE - ИМЯ ПРИЛАГАТЕЛЬНОЕ
   POSTPOSITION - ПОСЛЕЛОГ
   PRETEXT - ПРЕДЛОГ
   INFINITIVE - ИНФИНИТИВ
   TRANSGRESSIVE - ДЕЕПРИЧАСТИЕ
   PARTICIPLE - ПРИЧАСТИЕ
   PARTICLE - ЧАСТИЦА
   NUMERAL - ЧИСЛИТЕЛЬНОЕ
   CONJUNCTION - СОЮЗ
   MODALITY - ВВОДНЫЕ СЛОВА
 */

package sentence



import (
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"net/url"
	"Sima/config"
	"Sima/dialogflow"
)

var PythonNlpServer string


type Sentence struct {
	Value     string
	Entities   map[string]interface{}
}


type Sentences []Sentence


func  NewSentence(input string) *Sentence{
	S := new(Sentence)
	S.Value = input
	return S
}


func (S *Sentence) Prepare(marks map[string]string) {
	S.Value = strings.Replace(S.Value, marks["comma"], " ", -1)
}



func (S *Sentence) ToSpeechParts() {

	PythonNlpServer = config.Conf.Services.NLP.Server

	type Result struct {
		Result [][2]string `json:"result"`
	}

	Res := new(Result)

	data := url.Values{}
	data.Set("sentence", S.Value)


	client := &http.Client{}

	nlp_url := PythonNlpServer+"recognition/speech_part"

	r, _ := http.NewRequest("POST", nlp_url, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)

	if err != nil {panic(err)}

	json.NewDecoder(resp.Body).Decode(Res)

	fmt.Println(Res.Result)
}



func (S *Sentence) ToLanguageParts() {
	S.Entities = dialogflow.DetectIntentText(S.Value).Parameters
}

