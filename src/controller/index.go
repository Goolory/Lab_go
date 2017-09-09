package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"err_code": 0})
}
