package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var (
	rdb *redis.Client
)

func main_() {
	opt, err := redis.ParseURL("redis://default:72c6a1c60d694eb8b49f464368e9e7e0@apn1-immune-mongrel-34253.upstash.io:34253")
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)

	r := gin.Default()

	r.GET("/set/:key/:value", func(c *gin.Context) {
		key := c.Param("key")
		value := c.Param("value")
		err := rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "성공적으로 저장되었습니다."})
	})

	r.GET("/get/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, err := rdb.Get(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.Run(":8080")
}
