package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func CreateUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
