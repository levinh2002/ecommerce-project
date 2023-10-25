package main

import (
	"ecommerce-project/database"
	"ecommerce-project/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//app := controllers.NewApplication(database.ProductData(database.Cli, "Products"), database.UserData(database.Cli, "User"))
	app := database.Cli

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	//router.GET("/addtocart", app.AddToCart())
	//router.GET("/removeitem", app.RemoveItem())

	log.Fatal(router.Run(":" + port))
}
