package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	opt, err := redis.ParseURL("redis://default:72c6a1c60d694eb8b49f464368e9e7e0@apn1-immune-mongrel-34253.upstash.io:34253")
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)

	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := rdb.Set(ctx, user.Username, user.Password, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "성공적으로 회원 가입되었습니다"})
	})

	r.Run(":8080")
}
