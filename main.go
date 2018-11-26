package main

import (
	"time"
	"bufio"
	"Sima/database"
	"strconv"
	"Sima/train"
	"os"
	"Sima/tokens"
	"Sima/config"
	"fmt"
)

 var myscanner *bufio.Scanner

 var Config config.Config

func InitAll()  {
	Config = config.InitConfig()

	database.Connect()

	database.NewRedisConnect()

}



 var words  []string



func main() {

	InitAll()

	var S1 string

	//var Scores =  make(map[string]int)

	//var Soundex =  make(map[string]string)

	//file,_ := os.Open("RusS.txt")

	myscanner := bufio.NewScanner(os.Stdin) //os.StdIn


	S1 = ScanInput(myscanner)

	lang := "ru"

	TextModel := train.GetTrainingResult(lang)


	Text := tokens.NewText(S1,TextModel)

	_ = Text



	fmt.Println(Text.Sentences)
	time.Sleep(2000 * time.Second) // windows console window pause

}




func ScanInput(myscanner *bufio.Scanner)  string{
	myscanner.Scan()
	return myscanner.Text()
}



func is_numeric(Word string)  bool{
	if _, err := strconv.Atoi(Word); err == nil {
		return true
	}
	return false
}


















/* MAIN PART WAS HERE
	InitAll()

	var S1 string

	var Scores =  make(map[string]int)

	var Soundex =  make(map[string]string)

	file,_ := os.Open("RusS.txt")

	myscanner := bufio.NewScanner(file) //os.StdIn


	S1 = ScanInput(myscanner)

	S1 = abc.Replacer.Replace(S1)


	//fmt.Println(S1)


	Words := strings.Split(S1," ")

	fmt.Println(Words)



	for myscanner.Scan()  { Word := myscanner.Text()//for _,Word := range Words {

    	Input := Word

    	if len([]rune(Word)) < 2 && !is_numeric(Word) {continue}

		Word = rustemmer.GetWordBase(Word)


		Scores[Word] = abc.CountScore(Word)

		Soundex[Word] = phonetics.EncodeSoundex(unidecode.Unidecode(Word))


		var Word_arr =  make(map[string]string)
		Word_arr["word"] = Input
		Word_arr["base_word"] = Word
		Word_arr["soundex"] = Soundex[Word]
		Word_arr["score"] = strconv.Itoa(Scores[Word])
		InsertWord(Word_arr)

	}

   fmt.Println("END")
	//fmt.Println(Scores)
	//fmt.Println(Soundex)
 */