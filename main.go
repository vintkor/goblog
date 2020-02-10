package main

import (
	"goblog/news"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/news", news.ListView)
	r.Run()
}
