package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet)
	r.GET("/newsfeed", handler.Newsfeed_get(feed))
	r.POST("/newsfeed", handler.Newsfeed_post(feed))
	r.Run()
}
