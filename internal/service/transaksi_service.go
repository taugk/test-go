package service

import (
	"errors"
	"test-oldo/internal/model"
	"test-oldo/internal/repository"
)

type TransaksiService struct {
	trxRepo   *repository.TransaksiRepository
	paketRepo *repository.PaketRepository
	userRepo  *repository.UserRepository
}

func NewTransaksiService(
	trxRepo *repository.TransaksiRepository,
	paketRepo *repository.PaketRepository,
	userRepo *repository.UserRepository,
) *TransaksiService {
	return &TransaksiService{
		trxRepo:   trxRepo,
		paketRepo: paketRepo,
		userRepo:  userRepo,
	}
}

func (s *TransaksiService) Create(userID uint, paketID uint) (model.Transaksi, error) {

	_, err := s.userRepo.FindByID(userID)
	if err != nil {
		return model.Transaksi{}, errors.New("user tidak ditemukan")
	}

	paket, err := s.paketRepo.FindByID(paketID)
	if err != nil {
		return model.Transaksi{}, errors.New("paket tidak ditemukan")
	}

	trx := model.Transaksi{
		UserID:  userID,
		PaketID: paketID,
		Price:   paket.Price,
	}

	err = s.trxRepo.Create(&trx)
	return trx, err
}

func (s *TransaksiService) GetAll() ([]model.Transaksi, error) {
	return s.trxRepo.FindAll()
}

func (s *TransaksiService) GetByID(id uint) (model.Transaksi, error) {
	return s.trxRepo.FindByID(id)
}
