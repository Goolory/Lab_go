package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorAllowHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, Get")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-XSRF-Token, Authorization, JB-Filename")

	if origin := c.Request.Header.Get("Origin"); origin != "" {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
	if c.Request.Method == "OPTIONS" {
		c.Status(http.StatusOK)
	} else {
		c.Next()
	}
}
