package routes

import (
	"ecomercego/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductAdmin())
	incomingRoutes.GET("/users/productview", controllers.ViewProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProduct())
}
