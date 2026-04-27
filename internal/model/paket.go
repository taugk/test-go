package model

type Paket struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Quota        int    `json:"quota"`
	ActivePeriod int    `json:"active_period"`
}
