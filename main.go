package main

import (
	"ecomercego/controllers"
	"ecomercego/dao"
	"ecomercego/middleware"
	"ecomercego/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//init port server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//handler route
	app := controllers.NewApplication(dao.ProductData(dao.Client, "Products"), dao.UserData(dao.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Autentication())

	router.Get("/addtocart", app.AddToCart())
	router.GET("/removeitems", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instancebuy", app.InstantBuy())

	log.Fatal(router.Run(":", port))
}
