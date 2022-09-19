package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context
	}
}

func Login() gin.HandlerFunc {

}

func ProductAdmin() gin.HandlerFunc {

}

func ViewProduct() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}
