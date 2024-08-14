package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	mu   sync.RWMutex
	data map[interface{}]interface{}
}

func NewCache() *Cache {
	return &Cache{data: make(map[interface{}]interface{})}
}

func (c *Cache) Set(key string, user *User) {
	c.mu.Lock()
	c.data[key] = user
	c.mu.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if user, ok := c.data[key]; ok {
		return user
	}
	return nil
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {
	wg := sync.WaitGroup{}
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprintf("user-", i),
		})
	}

	time.Sleep(time.Second)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}

	wg.Wait()
}
