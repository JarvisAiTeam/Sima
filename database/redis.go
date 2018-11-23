package database


import (
  "github.com/go-redis/redis"
	"Sima/config"
)


var client *redis.Client




func NewRedisConnect() {
	params := config.Conf.Database.Redis

	client = redis.NewClient(&redis.Options{
		Addr:     params.Addr,
		Password: params.Password, // no password set
		DB:       params.DB,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	// Output: PONG <nil>
}


func GetRedisClient() *redis.Client{
	return client
}

func Exists(key string)  int{
	res, err := client.Exists(key).Result()
	if err != nil {
		panic(err)
	}
	return int(res)
}






