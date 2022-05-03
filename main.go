package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siriusol/configer/redis"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/keys", func(c *gin.Context) {
		keys, err := redis.New(&redis.Option{}).GetKeys()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"keys": keys,
		})
		return
	})

	r.Run(":8080")
}
