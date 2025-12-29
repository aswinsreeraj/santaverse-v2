package handlers

import (
	"encoding/json"
	"net/http"
	"santaverse/internal/database"
	"santaverse/internal/models"

	"github.com/gin-gonic/gin"
)

func GetGarage(c *gin.Context) {
	// Always use user_id = 1
	userID := 1

	// Query to fetch owned vehicles with mod status
	rows, err := database.DB.Query(`
		SELECT 
			ov.id, ov.item_id, i.name, i.type, i.category, ov.current_stats, i.image_url,
			(SELECT COUNT(*) FROM applied_mods am WHERE am.owned_vehicle_id = ov.id) > 0 as has_mods
		FROM owned_vehicles ov
		JOIN items i ON ov.item_id = i.id
		WHERE ov.user_id = $1
	`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch garage"})
		return
	}
	defer rows.Close()

	vehicles := []models.OwnedVehicle{} // Initialize as empty slice to return [] instead of null
	for rows.Next() {
		var v models.OwnedVehicle
		var currentStatsRaw []byte // Temporary variable for scanning current_stats
		if err := rows.Scan(&v.ID, &v.ItemID, &v.ItemName, &v.ItemType, &v.ItemCategory, &currentStatsRaw, &v.ItemImageURL, &v.HasMods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan vehicle data"})
			return
		}
		// Unmarshal current_stats from JSON byte array
		if err := json.Unmarshal(currentStatsRaw, &v.CurrentStats); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal vehicle stats"})
			return
		}
		vehicles = append(vehicles, v)
	}

	c.JSON(http.StatusOK, vehicles)
}

func BuyItem(c *gin.Context) {
	var req struct {
		ItemID int `json:"item_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Fetch base stats of the item
	var baseStats []byte
	err := database.DB.QueryRow("SELECT stats FROM items WHERE id = $1", req.ItemID).Scan(&baseStats)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// user_id = 1
	userID := 1
	_, err = database.DB.Exec("INSERT INTO owned_vehicles (item_id, user_id, current_stats) VALUES ($1, $2, $3)", req.ItemID, userID, baseStats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to purchase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase successful"})
}

func ApplyMod(c *gin.Context) {
	var req struct {
		VehicleID int `json:"vehicle_id"`
		ModID     int `json:"mod_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 1. Get current stats of vehicle
	var currentStatsRaw []byte
	err := database.DB.QueryRow("SELECT current_stats FROM owned_vehicles WHERE id = $1", req.VehicleID).Scan(&currentStatsRaw)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	var currentStats models.Stats
	json.Unmarshal(currentStatsRaw, &currentStats)

	// 2. Get mod stats
	var modStatsRaw []byte
	err = database.DB.QueryRow("SELECT stat_boosts FROM mods WHERE id = $1", req.ModID).Scan(&modStatsRaw)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mod not found"})
		return
	}

	var modBoosts models.Stats // Using existing Stats struct assuming fields match "stat_boosts"
	json.Unmarshal(modStatsRaw, &modBoosts)

	// 3. Apply mod (Simple addition)
	currentStats.Speed += modBoosts.Speed
	currentStats.Stamina += modBoosts.Stamina
	currentStats.Capacity += modBoosts.Capacity
	currentStats.Magic += modBoosts.Magic

	// 4. Update vehicle stats
	newStatsRaw, _ := json.Marshal(currentStats)
	_, err = database.DB.Exec("UPDATE owned_vehicles SET current_stats = $1 WHERE id = $2", newStatsRaw, req.VehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stats"})
		return
	}

	// 5. Record applied mod
	_, err = database.DB.Exec("INSERT INTO applied_mods (owned_vehicle_id, mod_id) VALUES ($1, $2)", req.VehicleID, req.ModID)
	if err != nil {
		// Log error but don't fail the request since stats are updated (simplification)
		// actually, ideally generic error handling.
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mod applied successfully", "new_stats": currentStats})
}

func SellVehicle(c *gin.Context) {
	var req struct {
		VehicleID int `json:"vehicle_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Delete vehicle (Cascade should handle applied_mods if configured, else manual delete)
	// For safety, let's delete applied_mods first
	_, err := database.DB.Exec("DELETE FROM applied_mods WHERE owned_vehicle_id = $1", req.VehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove mods"})
		return
	}

	_, err = database.DB.Exec("DELETE FROM owned_vehicles WHERE id = $1", req.VehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sell vehicle"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vehicle returned to Santa"})
}

func ResetVehicle(c *gin.Context) {
	var req struct {
		VehicleID int `json:"vehicle_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 1. Get original item ID
	var itemID int
	err := database.DB.QueryRow("SELECT item_id FROM owned_vehicles WHERE id = $1", req.VehicleID).Scan(&itemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	// 2. Fetch base stats
	var baseStats []byte
	err = database.DB.QueryRow("SELECT stats FROM items WHERE id = $1", itemID).Scan(&baseStats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch base stats"})
		return
	}

	// 3. Reset stats in owned_vehicles
	_, err = database.DB.Exec("UPDATE owned_vehicles SET current_stats = $1 WHERE id = $2", baseStats, req.VehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revert stats"})
		return
	}

	// 4. Remove all applied mods
	_, err = database.DB.Exec("DELETE FROM applied_mods WHERE owned_vehicle_id = $1", req.VehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove mods"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vehicle reset to factory settings"})
}
