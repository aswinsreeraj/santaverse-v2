package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"santaverse/internal/database"
	"santaverse/internal/handlers"
)

func main() {
	database.Connect()
	database.InitSchema("./internal/database/schema.sql")

	r := gin.Default()

	// Serve static files
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")
	r.Static("/img", "./web/img")
	r.LoadHTMLGlob("web/*.html")

	// Routes
	api := r.Group("/api")
	{
		api.GET("/marketplace", handlers.GetMarketplaceItems)
		api.GET("/marketplace/:id", handlers.GetItemDetails)
		api.GET("/garage", handlers.GetGarage)
		api.POST("/buy", handlers.BuyItem)
		api.GET("/mods", handlers.GetMods)
		api.POST("/garage/mod", handlers.ApplyMod)
		api.POST("/garage/sell", handlers.SellVehicle)
		api.POST("/garage/reset", handlers.ResetVehicle)
	}

	// HTML Pages Routes (Simple serving of HTML files, frontend handles data fetching) - or we can keep serving static HTMLs generally
	// The static serving above `r.LoadHTMLGlob("web/*.html")` covers serving files if we render them server side or if we just serve the template.
	// Since we are doing SPA-like with vanilla JS, we might want to serve the specific html files at specific paths.

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Santaverse Garage"})
	})
	r.GET("/product", func(c *gin.Context) {
		c.HTML(http.StatusOK, "product.html", gin.H{"title": "Product Details"})
	})
	r.GET("/garage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "garage.html", gin.H{"title": "My Garage"})
	})

	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
