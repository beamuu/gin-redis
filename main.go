package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func get(client *redis.Client, key string) (string, error) {
	return client.Get(key).Result()
}

func set(client *redis.Client, key string, value interface{}) error {
	return client.Set(key, value, 0).Err()
}
func Get(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("key")
		value, err := get(client, query)
		if err == nil {
			c.JSON(404, gin.H{"message": "Error"})
		} else {
			c.JSON(200, gin.H{"message": "success", "key": query, "value": value})
		}
	}
}

func Set(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")
		value := c.Query("value")

		err := set(client, key, value)
		if err == nil {
			c.JSON(404, gin.H{"message": "Error"})
		} else {
			c.JSON(200, gin.H{"message": "success", "key": key, "value": value})
		}
	}
}

func main() {

	// Redis intialize

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	r := gin.Default()
	r.Use()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/set", func(c *gin.Context) {
		key := c.Query("key")
		value := c.Query("value")

		err := set(client, key, value)
		if err != nil {
			c.JSON(404, gin.H{"message": "Error"})
		} else {
			c.JSON(200, gin.H{"message": "success", "key": key, "value": value})
		}
	})

	r.GET("/get", func(c *gin.Context) {
		query := c.Query("key")
		value, err := get(client, query)
		if err != nil {
			c.JSON(404, gin.H{"message": "Error"})
		} else {
			c.JSON(200, gin.H{"message": "success", "key": query, "value": value})
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
