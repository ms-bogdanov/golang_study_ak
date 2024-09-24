package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Cacher interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cacher {
	return &Cache{
		client: client,
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	return c.client.Set(key, value, 0).Err()
}

func (c *Cache) Get(key string) (interface{}, error) {
	return c.client.Get(key).Result()
}

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	cache := NewCache(client)

	err := cache.Set("some:key", "value")
	if err != nil {
		panic(err)
	}

	value, err := cache.Get("some:key")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)

	user := &User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	err = cache.Set(fmt.Sprintf("user:%v", user.ID), user)
	if err != nil {
		panic(err)
	}

	value, err = cache.Get("key")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
