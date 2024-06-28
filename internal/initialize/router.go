package initialize

import (
	"net/http"

	"github.com/duonglc27092000/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1/2024")
	{
		v1.GET("ping", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserByID)
		// v1.PUT("ping", Pong)
		// v1.PATCH("ping", Pong)
		// v1.DELETE("ping", Pong)
		// v1.HEAD("ping", Pong)
		// v1.OPTIONS("ping", Pong)
	}

	// v2 := r.Group("v2/2024")
	// {
	// 	v2.GET("ping", Pong)
	// 	v2.PUT("ping", Pong)
	// 	v2.PATCH("ping", Pong)
	// 	v2.DELETE("ping", Pong)
	// 	v2.HEAD("ping", Pong)
	// 	v2.OPTIONS("ping", Pong)
	// }
	return r
}
func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "LuongDuong")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"d1", "d2", "d3"},
	})
}
