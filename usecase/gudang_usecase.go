package usecase

import (
	"barang/model"
	"barang/repository"
)

type GudangUsecase interface { // layer untuk komunikasi | jembatan antar layer
	RegisterGudang(newGudang model.Gudang) error
    FindAllGudang() ([]model.Gudang, error)
    UpdateGudang(newGudang model.Gudang) error
    DeleteGudangById(id string) error
    
}

type gudangUsecase struct {
    repo repository.GudangRepository
}

func (g *gudangUsecase) RegisterGudang(newGudang model.Gudang) error {
    return g.repo.InsertGudang(newGudang)
}

func (g *gudangUsecase) FindAllGudang() ([]model.Gudang, error) {
    return g.repo.GetListGudang()
}

func (g *gudangUsecase) UpdateGudang(newGudang model.Gudang) error {    
    return g.repo.UpdateGudang(newGudang)
}

func (g *gudangUsecase) DeleteGudangById(id string) error {
    return g.repo.DeleteGudangById(id)
}

func NewGudangUsecase(repo repository.GudangRepository) GudangUsecase {
    return &gudangUsecase{
        repo: repo,
    }
}