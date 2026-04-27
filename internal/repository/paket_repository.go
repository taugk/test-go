package repository

import (
	"test-oldo/internal/model"

	"gorm.io/gorm"
)

type PaketRepository struct {
	DB *gorm.DB
}

func NewPaketRepository(db *gorm.DB) *PaketRepository {
	return &PaketRepository{DB: db}
}

func (r *PaketRepository) Create(paket *model.Paket) error {
	return r.DB.Create(paket).Error
}

func (r *PaketRepository) FindAll() ([]model.Paket, error) {
	var pakets []model.Paket
	err := r.DB.Find(&pakets).Error
	return pakets, err
}

func (r *PaketRepository) FindByID(id uint) (model.Paket, error) {
	var paket model.Paket
	err := r.DB.First(&paket, id).Error
	return paket, err
}

func (r *PaketRepository) Update(paket *model.Paket) error {
	return r.DB.Save(paket).Error
}

func (r *PaketRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Paket{}, id).Error
}
