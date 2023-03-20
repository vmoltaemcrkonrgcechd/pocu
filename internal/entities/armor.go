package entities

type Armor struct {
	ID uint16 `json:"id"`
	ArmorDTO
}

type ArmorDTO struct {
	Name       string  `json:"name"`
	Protection float32 `json:"protection"`
	Weight     float32 `json:"weight"`
}

type AllArmorDTO struct {
	Armor         []Armor `json:"armor"`
	Quantity      uint    `json:"quantity"`
	MinWeight     float32 `json:"minWeight"`
	MaxWeight     float32 `json:"maxWeight"`
	MinProtection float32 `json:"minProtection"`
	MaxProtection float32 `json:"maxProtection"`
}
