package usecase

import (
	"barang/model"
	"barang/repository"
)

type BarangUsecase interface { // layer untuk komunikasi | jembatan antar layer
	RegisterBarang(newBarang model.Barang) error
    FindAllBarang() ([]model.Barang, error)
    UpdateBarang(newBarang model.Barang) error
    DeleteBarangById(id string) error
    
}

type barangUsecase struct {
    repo repository.BarangRepository
}

func (b *barangUsecase) RegisterBarang(newBarang model.Barang) error {
    return b.repo.InsertBarang(newBarang)
}

func (b *barangUsecase) FindAllBarang() ([]model.Barang, error) {
    return b.repo.GetListBarang()
}

func (b *barangUsecase) UpdateBarang(newBarang model.Barang) error {    
    return b.repo.UpdateBarang(newBarang)
}

func (b *barangUsecase) DeleteBarangById(id string) error {
    return b.repo.DeleteBarangById(id)
}

func NewBarangUsecase(repo repository.BarangRepository) BarangUsecase {
    return &barangUsecase{
        repo: repo,
    }
}