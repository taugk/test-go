package service

import (
	"errors"
	"test-oldo/internal/model"
	"test-oldo/internal/repository"
)

type PaketService struct {
	repo *repository.PaketRepository
}

func NewPaketService(repo *repository.PaketRepository) *PaketService {
	return &PaketService{repo: repo}
}

func (s *PaketService) Create(paket *model.Paket) error {
	if paket.Name == "" {
		return errors.New("nama paket tidak boleh kosong")
	}
	if paket.Price <= 0 {
		return errors.New("harga harus lebih dari 0")
	}
	if paket.Quota <= 0 {
		return errors.New("quota harus lebih dari 0")
	}
	if paket.ActivePeriod <= 0 {
		return errors.New("masa aktif harus lebih dari 0")
	}

	return s.repo.Create(paket)
}

func (s *PaketService) GetAll() ([]model.Paket, error) {
	return s.repo.FindAll()
}

func (s *PaketService) GetByID(id uint) (model.Paket, error) {
	return s.repo.FindByID(id)
}

func (s *PaketService) Update(id uint, input *model.Paket) (model.Paket, error) {
	paket, err := s.repo.FindByID(id)
	if err != nil {
		return paket, err
	}

	paket.Name = input.Name
	paket.Price = input.Price
	paket.Quota = input.Quota
	paket.ActivePeriod = input.ActivePeriod

	err = s.repo.Update(&paket)
	return paket, err
}

func (s *PaketService) Delete(id uint) error {
	return s.repo.Delete(id)
}
