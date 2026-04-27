package model

import "time"

type Transaksi struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	PaketID   uint      `json:"paket_id" gorm:"not null"`
	Price     int       `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`

	User  *User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Paket *Paket `gorm:"foreignKey:PaketID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
