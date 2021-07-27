package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func  getRedishData(myquery string) string {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	val, err := client.Get(myquery).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val

}


func  setRedishData(mydata string, myquery string) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	var err = client.Set(myquery, mydata, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}
