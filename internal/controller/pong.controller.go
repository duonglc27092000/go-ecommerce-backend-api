package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type PongController struct{}

func NewPongController() *PongController{
	return &PongController{}
}

func (p*PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "LuongDuong")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"d1", "d2", "d3"},
	})
}