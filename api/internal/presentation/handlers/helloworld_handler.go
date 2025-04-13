package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "helloworld")
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/helloworld", HelloWorldHandler)
}
