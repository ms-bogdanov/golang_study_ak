package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mu sync.RWMutex

type User struct {
	ID   int
	Name string
}

type Cache struct {
	cache map[string]*User
}

func NewCache() *Cache {
	data := make(map[string]*User)
	return &Cache{cache: data}
}

func (c *Cache) Set(key string, user *User) {
	mu.Lock()
	c.cache[key] = user
	mu.Unlock()
}

func (c *Cache) Get(key string) *User {
	mu.RLock()
	result := c.cache[key]
	mu.RUnlock()
	return result
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		j := i

		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   j,
			Name: fmt.Sprint("user-", j),
		})
	}

	time.Sleep(5 * time.Millisecond)

	for i := 0; i < 100; i++ {
		j := i

		go func(j int) {
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(j))))
		}(j)
	}

	time.Sleep(5 * time.Millisecond)
}
