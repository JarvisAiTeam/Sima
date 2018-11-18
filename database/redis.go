package database


import (
  "github.com/go-redis/redis"
)


var client *redis.Client




func NewRedisConnect() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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






