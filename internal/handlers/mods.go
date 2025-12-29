package handlers

import (
	"net/http"
	"santaverse/internal/database"
	"santaverse/internal/models"

	"github.com/gin-gonic/gin"
)

func GetMods(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, price, stat_boosts FROM mods")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var modsList []models.Mod
	for rows.Next() {
		var m models.Mod
		var statsData []byte
		if err := rows.Scan(&m.ID, &m.Name, &m.Price, &statsData); err != nil {
			continue
		}
		m.StatBoosts = statsData
		modsList = append(modsList, m)
	}

	c.JSON(http.StatusOK, modsList)
}
