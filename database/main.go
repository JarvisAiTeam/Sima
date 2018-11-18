package database


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)


var db *sql.DB





func params() string {

	info := fmt.Sprintf("Nick01:password@/Words_info_bd")
	return info
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func Connect() {
	var err error
	db, err = sql.Open("mysql",params())
	fatal(err)



}






func CheckWord(Word map[string]string) (id int)  {
	row := db.QueryRow("SELECT `*` FROM `words` WHERE ``=?",Word["word"])
	row.Scan(&id)
	return id
}

func InsertWord(Word map[string]string)  (sql.Result, error){
	return db.Exec("INSERT INTO `words` (`word`,`base_word`,`soundexed` ,`score`) VALUES (?, ?, ?, ?)",
		Word["word"],Word["base_word"], Word["soundex"], Word["score"])
}






