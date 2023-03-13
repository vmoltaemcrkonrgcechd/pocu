package entities

type Weapon struct {
	ID uint16 `json:"id"`
	WeaponDTO
}

type WeaponDTO struct {
	Name   string  `json:"name"`
	Attack float32 `json:"attack"`
	Weight float32 `json:"weight"`
}

type AllWeaponsDTO struct {
	Weapons   []Weapon `json:"weapons"`
	Quantity  uint     `json:"quantity"`
	MinAttack float32  `json:"minAttack"`
	MaxAttack float32  `json:"maxAttack"`
	MinWeight float32  `json:"minWeight"`
	MaxWeight float32  `json:"maxWeight"`
}
