package models

import "encoding/json"

type Stats struct {
	Speed    int `json:"speed"`
	Stamina  int `json:"stamina"`
	Capacity int `json:"capacity"`
	Magic    int `json:"magic"`
}

type Item struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Category string          `json:"category"`
	Price    int             `json:"price"`
	ImageURL string          `json:"image_url"`
	Stats    json.RawMessage `json:"stats"` // store as raw JSON for flexibility or use Stats struct if consistent
}

// Helper to parse stats if needed, but for now we might send raw json to frontend
// Or we can unmarshal into Stats struct. Let's use specific struct for better type safety.

type ItemWithStats struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	ImageURL string `json:"image_url"`
	Stats    Stats  `json:"stats"`
}

type Mod struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Price      int             `json:"price"`
	StatBoosts json.RawMessage `json:"stat_boosts"`
}

type OwnedVehicle struct {
	ID           int             `json:"id"`
	ItemID       int             `json:"item_id"`
	UserID       int             `json:"user_id"`
	CurrentStats json.RawMessage `json:"current_stats"`
	// Join fields for display
	ItemName     string `json:"item_name,omitempty"`
	ItemImageURL string `json:"item_image_url,omitempty"`
	ItemType     string `json:"item_type,omitempty"`
	ItemCategory string `json:"item_category,omitempty"`
	HasMods      bool   `json:"has_mods"`
}
