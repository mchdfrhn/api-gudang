package repository

import (
	"barang/model"
	"barang/utils"
	"database/sql"
)

type GudangRepository interface {
	GetListGudang() ([]model.Gudang, error)
	InsertGudang(newGudang model.Gudang) error
	UpdateGudang(newGudang model.Gudang) error
	DeleteGudangById(id string) error
}

type gudangRepository struct {
	db *sql.DB
}

func (r *gudangRepository) GetListGudang() ([]model.Gudang, error) {
	rows, err := r.db.Query(utils.READ_GUDANG)
	if err != nil {
		return nil, err
	}
	var gudangs []model.Gudang
	for rows.Next() {
		var gudang model.Gudang
		err = rows.Scan(
			&gudang.Kode,
			&gudang.Nama,
		)
		if err != nil {
			return nil, err
		}
		gudangs = append(gudangs, gudang)
	}

	return gudangs, nil
}

func (r *gudangRepository) InsertGudang(newGudang model.Gudang) error {
	_, err := r.db.Exec(utils.CREATE_GUDANG,
		newGudang.Kode,
		newGudang.Nama,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *gudangRepository) UpdateGudang(newGudang model.Gudang) error {
	_, err := r.db.Exec(utils.UPDATE_GUDANG,
		newGudang.Nama,
		newGudang.Kode,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *gudangRepository) DeleteGudangById(id string) error {
	_, err := r.db.Exec(utils.DELETE_GUDANG, id)
	if err != nil {
		return err
	}
	return nil
}

func NewGudangRepository(db *sql.DB) GudangRepository {
	return &gudangRepository{
		db: db,
	}
}