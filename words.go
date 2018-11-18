package main

import (
	"Sima/database"
	"fmt"
)



func CheckIfExist(word map[string]string) {
	database.CheckWord(word)
}


func InsertWord(word map[string]string) {
	_, err := database.InsertWord(word)
	fmt.Println(err)
}
