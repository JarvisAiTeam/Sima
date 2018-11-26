package config

import (
	"github.com/BurntSushi/toml"
)


type Config struct {
	Database Database
	Services Services
}

type Database struct {
	MySQL MySQL
	Redis Redis
}

type MySQL struct {
	User string
	Password string
	DatabaseName string
}

type Redis struct {
	Addr string
	Password string
	DB  int
}

type Services struct {
	NLP NLP
	Dialogflow Dialogflow
}

type NLP struct {
	Server string
}

type Dialogflow struct {
	Token string
}



var Conf Config


func InitConfig() Config{
	if _, err := toml.DecodeFile("./config.toml", &Conf); err != nil {
		// handle error
		panic(err)
	}
	return Conf
}


