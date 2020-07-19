package main

import (
  "github.com/gomodule/redigo/redis"
  "fmt"
)

func main(){

  c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	res, err := redis.Strings(c.Do("HKEYS", "people1"))
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

  fmt.Println("res:  ",res[0])

	username, err := redis.String(c.Do("HGET", "people2","name"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

}
