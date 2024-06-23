package repository

import (
	"barang/model"
	"barang/utils"
	"database/sql"
)

type BarangRepository interface {
	GetListBarang() ([]model.Barang, error)
	InsertBarang(newBarang model.Barang) error
	UpdateBarang(newBarang model.Barang) error
	DeleteBarangById(id string) error
}

type barangRepository struct {
	db *sql.DB
}

func (r *barangRepository) GetListBarang() ([]model.Barang, error) {
	rows, err := r.db.Query(utils.READ_BARANG)
	if err != nil {
		return nil, err
	}
	var barangs []model.Barang
	for rows.Next() {
		var barang model.Barang
		err = rows.Scan(
			&barang.Kode,
			&barang.Nama,
			&barang.Harga,
			&barang.Jumlah,
			&barang.Expired,
			&barang.Gudang.Kode,
			&barang.Gudang.Nama,
		)
		if err != nil {
			return nil, err
		}
		barangs = append(barangs, barang)
	}

	return barangs, nil
}

func (r *barangRepository) InsertBarang(newBarang model.Barang) error {
	_, err := r.db.Exec(utils.CREATE_BARANG,
		newBarang.Kode,
		newBarang.Nama,
		newBarang.Harga,
		newBarang.Jumlah,
		newBarang.Expired,
		newBarang.Gudang.Kode,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) UpdateBarang(newBarang model.Barang) error {
	_, err := r.db.Exec(utils.UPDATE_BARANG,
		newBarang.Nama,
		newBarang.Harga,
		newBarang.Jumlah,
		newBarang.Expired,
		newBarang.Kode,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) DeleteBarangById(id string) error {
	_, err := r.db.Exec(utils.DELETE_BARANG, id)
	if err != nil {
		return err
	}
	return nil
}

func NewBarangRepository(db *sql.DB) BarangRepository {
	return &barangRepository{
		db: db,
	}
}