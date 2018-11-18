package recognizer

import (
	"net/http"
	"encoding/json"
	"strings"
	"net/url"
)


type LangResponse struct {
	Code int `json:"code"`
	Lang string `json:"lang"`
}




type Recognizer struct {

}


func GetLang(Input string) string{

	return requestLang(Input)
}



func requestLang(Input string) string{

	langResp := new(LangResponse)

	data := url.Values{}
	data.Set("text", Input)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", "https://translate.yandex.net/api/v1.5/tr.json/detect?hint=&key=trnsl.1.1.20181025T183755Z.dbec268dad6e0333.df36bb98b7fedcc1c620ca43f906070d1b752d3b", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)

	if err != nil {panic(err)}

	json.NewDecoder(resp.Body).Decode(langResp)

	return langResp.Lang

}






