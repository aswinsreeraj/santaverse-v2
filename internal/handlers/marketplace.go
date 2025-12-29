package handlers

import (
	"log"
	"net/http"
	"santaverse/internal/database"
	"santaverse/internal/models"

	"github.com/gin-gonic/gin"
)

func GetMarketplaceItems(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, type, category, price, image_url, stats FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var i models.Item
		// For JSONB, we can scan into []byte or json.RawMessage
		var statsData []byte
		if err := rows.Scan(&i.ID, &i.Name, &i.Type, &i.Category, &i.Price, &i.ImageURL, &statsData); err != nil {
			log.Println("Error scanning item:", err)
			continue
		}
		i.Stats = statsData
		items = append(items, i)
	}

	c.JSON(http.StatusOK, items)
}

func GetItemDetails(c *gin.Context) {
	id := c.Param("id")
	var i models.Item
	var statsData []byte

	err := database.DB.QueryRow("SELECT id, name, type, category, price, image_url, stats FROM items WHERE id = $1", id).Scan(
		&i.ID, &i.Name, &i.Type, &i.Category, &i.Price, &i.ImageURL, &statsData,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	i.Stats = statsData
	c.JSON(http.StatusOK, i)
}
