package handler

import (
	"net/http"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func Newsfeed_get(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := feed.GetAll()
		c.JSON(http.StatusOK, res)
	}

}
