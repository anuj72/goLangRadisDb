package redis

import (
	"RedisProject/config"
	"fmt"
	"github.com/go-redis/redis"
)

var add string = config.NewRepository().RedisConnectionURL()
var client = redis.NewClient(&redis.Options{
	Addr:     add,
	Password: "",
	DB:       0,
})

func GetRedisData(myquery string) string {
	val, err := client.Get(myquery).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func SetRedisData(mydata string, myquery string) {
	var err = client.Set(myquery, mydata, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}
