package repository

import (
	"test-oldo/internal/model"

	"gorm.io/gorm"
)

type TransaksiRepository struct {
	DB *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	return &TransaksiRepository{DB: db}
}

func (r *TransaksiRepository) Create(trx *model.Transaksi) error {
	return r.DB.Create(trx).Error
}

func (r *TransaksiRepository) FindAll() ([]model.Transaksi, error) {
	var data []model.Transaksi
	err := r.DB.Preload("User").Preload("Paket").Find(&data).Error
	return data, err
}

func (r *TransaksiRepository) FindByID(id uint) (model.Transaksi, error) {
	var trx model.Transaksi
	err := r.DB.Preload("User").Preload("Paket").First(&trx, id).Error
	return trx, err
}
